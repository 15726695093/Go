package service

import (
	v1 "clock-in/api/worktime/service/v1"
	"clock-in/app/worktime/service/internal/biz"
	"clock-in/app/worktime/service/internal/data/ent"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService
type WorkTimeService struct {
	v1.UnimplementedWorktimeServiceServer

	uc  *biz.WorkTimeUsecase
	log *log.Helper
}

var _ v1.WorktimeServiceServer = (*WorkTimeService)(nil)

func NewWorkTimeService(uc *biz.WorkTimeUsecase, logger log.Logger) *WorkTimeService {
	return &WorkTimeService{uc: uc, log: log.NewHelper(logger)}
}

func (svc *WorkTimeService) GetUserWorkTime(ctx context.Context, in *v1.GetUserWorkTimeRequest) (reply *v1.GetUserWorkTimeReply, err error) {
	worktime, err := svc.uc.GetUserWorkTime(ctx, in.User, in.Day)
	switch {
	case err != nil:
		err = v1.ErrorUnknownError(err.Error())
	case len(worktime) == 0:
		err = v1.ErrorRecordExisted("error: work time record not existed")
	}
	if err == nil {
		reply = &v1.GetUserWorkTimeReply{
			Worktime: bulk2ProtoWorkTime(worktime),
		}
	}
	return
}
func (svc *WorkTimeService) CreateWorkTime(ctx context.Context, in *v1.CreateWorkTimeRequest) (reply *v1.CreateWorkTimeReply, err error) {
	err = svc.uc.CreateWorkTime(ctx, in.User, bulk2BizRecord(in.Record))
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = v1.ErrorRecordExisted("error: worktime record existed")
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	reply = &v1.CreateWorkTimeReply{}
	return
}

func bulk2ProtoWorkTime(source []*biz.WorkTime) []*v1.WorkTime {
	worktime := make([]*v1.WorkTime, 0, len(source))
	for _, v := range source {
		worktime = append(worktime, &v1.WorkTime{
			Id:     v.Id,
			Minute: v.Minute,
			Day:    v.Day,
		})
	}
	return worktime
}

func bulk2BizRecord(source []*v1.Record) []*biz.Record {
	record := make([]*biz.Record, 0, len(source))
	for _, v := range source {
		record = append(record, &biz.Record{
			Day:    v.Day,
			Moment: v.Moment,
			User:   v.User,
		})
	}
	return record
}
