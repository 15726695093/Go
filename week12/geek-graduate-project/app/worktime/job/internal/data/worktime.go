package data

import (
	"clock-in/app/worktime/job/internal/biz"
	"context"

	worktimev1 "clock-in/api/worktime/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type workTimeRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.WorkTimeRepo = (*workTimeRepo)(nil)

// NewUserRepo .
func NewWorkTimeRepo(data *Data, logger log.Logger) biz.WorkTimeRepo {
	return &workTimeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *workTimeRepo) CaculateWorktime(ctx context.Context, records *biz.Record) error {
	_, err := repo.data.wc.CreateWorkTime(ctx, &worktimev1.CreateWorkTimeRequest{
		User: records.User,
		Record: []*worktimev1.Record{
			{Day: records.Day1, Moment: records.Moment1},
			{Day: records.Day2, Moment: records.Moment2},
		},
	})
	return err
}
