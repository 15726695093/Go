package data

import (
	recordv1 "clock-in/api/record/v1"
	"clock-in/app/clockin/interface/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.RecordRepo = (*recordRepo)(nil)

// NewUserRepo .
func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *recordRepo) ClockInOnWork(ctx context.Context) error {
	_, err := r.data.rc.ClockInOnWork(ctx, &recordv1.ClockInOnWorkRequest{
		User: int64(ctx.Value("user").(int)),
	})
	return err
}

func (r *recordRepo) ClockInOffWork(ctx context.Context) error {
	_, err := r.data.rc.ClockInOffWork(ctx, &recordv1.ClockInOffWorkRequest{
		User: int64(ctx.Value("user").(int)),
	})
	return err
}
