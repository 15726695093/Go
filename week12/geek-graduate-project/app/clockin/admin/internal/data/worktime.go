package data

import (
	worktimev1 "clock-in/api/worktime/service/v1"
	"clock-in/app/clockin/admin/internal/biz"
	"context"

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

func (r *workTimeRepo) GetUserWorkTime(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
	reply, err := r.data.wc.GetUserWorkTime(ctx, &worktimev1.GetUserWorkTimeRequest{
		User: user,
		Day:  day,
	})
	if err != nil {
		return nil, err
	}
	return bulk2BizWorkTime(reply.Worktime), nil
}

func bulk2BizWorkTime(source []*worktimev1.WorkTime) []*biz.WorkTime {
	worktime := make([]*biz.WorkTime, 0, len(source))
	for _, v := range source {
		worktime = append(worktime, &biz.WorkTime{
			Id:     v.Id,
			Day:    v.Day,
			Minute: v.Minute,
		})
	}
	return worktime
}
