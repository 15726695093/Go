package data

import (
	"clock-in/app/worktime/service/internal/biz"
	"clock-in/app/worktime/service/internal/data/ent"
	"clock-in/app/worktime/service/internal/data/ent/worktime"
	"context"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type workTimeRepo struct {
	data *Data
	log  *log.Helper
}

var ErrWorkTimeNotFoundInCache = errors.New("error: work time not existed in cache")

var _ biz.WorkTimeRepo = (*workTimeRepo)(nil)

// NewUserRepo .
func NewWorkTimeRepo(data *Data, logger log.Logger) biz.WorkTimeRepo {
	return &workTimeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *workTimeRepo) GetUserWorkTime(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
	reply, err := r.getUserWorkTimeFromCache(ctx, user, day)
	if err != nil {
		reply, err = r.getUserWorkTimeFromDB(ctx, user, day)
		if err != nil {
			return nil, err
		}
		return reply, r.storeUserWorkTime2Cache(ctx, user, reply)
	}
	return reply, nil
}

func (r *workTimeRepo) getUserWorkTimeFromDB(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
	rows, err := r.data.db.Worktime.Query().
		Where(
			worktime.UserEQ(user),
			worktime.DayIn(day...),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return bulk2zBizWorkTime(rows), nil
}

func (r *workTimeRepo) getUserWorkTimeFromCache(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
	val, err := r.data.rc.HGetAll(ctx, strconv.Itoa(int(user))).Result()
	if err != nil {
		return nil, err
	}
	worktime := make([]*biz.WorkTime, 0)
	for _, v := range day {
		s, ok := val[strconv.Itoa(int(v))]
		if !ok {
			return nil, ErrWorkTimeNotFoundInCache
		}
		minute, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		worktime = append(worktime, &biz.WorkTime{
			Day:    v,
			Minute: minute,
		})
	}
	return worktime, nil
}

func (r *workTimeRepo) storeUserWorkTime2Cache(ctx context.Context, user int64, record []*biz.WorkTime) error {
	data := make(map[string]interface{})
	for _, v := range record {
		data[strconv.Itoa(int(v.Day))] = v.Minute
	}
	return r.data.rc.HSet(ctx, strconv.Itoa(int(user)), data).Err()
}

func (r *workTimeRepo) CreateWorkTime(ctx context.Context, user int64, records []*biz.Record) error {
	gap := math.Abs(float64(records[0].Moment)-float64(records[1].Moment)) / 60
	_, err := r.data.db.Worktime.Create().
		SetDay(records[0].Day).
		SetUser(user).
		SetMinute(int64(gap)).
		SetCreatedAt(time.Now()).
		Save(ctx)
	return err
}

func toBizWorkTime(source *ent.Worktime) *biz.WorkTime {
	return &biz.WorkTime{
		Id:     int64(source.ID),
		Day:    source.Day,
		Minute: source.Minute,
	}
}

func bulk2zBizWorkTime(source []*ent.Worktime) []*biz.WorkTime {
	worktime := make([]*biz.WorkTime, 0, len(source))
	for _, v := range source {
		worktime = append(worktime, toBizWorkTime(v))
	}
	return worktime
}
