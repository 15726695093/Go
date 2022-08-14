package data

import (
	"clock-in/app/record/service/internal/biz"
	"clock-in/app/record/service/internal/data/ent"
	"clock-in/app/record/service/internal/data/ent/record"
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/streadway/amqp"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

var ErrRecordNotComplete = errors.New("error: user clockin records does not complete")

var _ biz.RecordRepo = (*recordRepo)(nil)

// NewUserRepo .
func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *recordRepo) ClockIn(ctx context.Context, record *biz.Record) error {
	_, err := r.data.db.Record.Create().
		SetUser(record.User).
		SetDay(record.Day).
		SetType(record.Type).
		SetCreatedAt(record.Clock).
		Save(ctx)
	return err
}

func (r *recordRepo) CaculateWorkTime(ctx context.Context, user int64, day int64) error {
	rows, err := r.data.db.Record.Query().
		Where(
			record.UserEQ(user),
			record.DayEQ(day),
		).
		All(ctx)
	if err != nil {
		return err
	}
	message, err := buildMessage(rows)
	if err != nil {
		return err
	}
	return r.data.mq.Publish(
		"",         // exchange
		"worktime", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
}

func buildMessage(record []*ent.Record) ([]byte, error) {
	if len(record) < 2 {
		return nil, ErrRecordNotComplete
	}
	message := strconv.Itoa(int(record[0].User))
	for i := 0; i < 2; i++ {
		message = message + "-" + strconv.Itoa(int(record[i].Day)) + "-" + strconv.Itoa(int(record[i].CreatedAt.Unix()))
	}
	return []byte(message), nil
}
