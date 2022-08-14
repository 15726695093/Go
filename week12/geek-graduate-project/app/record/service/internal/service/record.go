package service

import (
	v1 "clock-in/api/record/v1"
	"clock-in/app/record/service/internal/biz"
	"clock-in/app/record/service/internal/data/ent"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService
type RecordService struct {
	v1.UnimplementedRecordServiceServer

	uc  *biz.RecordUsecase
	log *log.Helper
}

var _ v1.RecordServiceServer = (*RecordService)(nil)

// NewUserService new a greeter service.
func NewUserService(uc *biz.RecordUsecase, logger log.Logger) *RecordService {
	return &RecordService{uc: uc, log: log.NewHelper(logger)}
}

func (svc *RecordService) ClockInOnWork(ctx context.Context, in *v1.ClockInOnWorkRequest) (*v1.ClockInOnWorkReply, error) {
	err := svc.uc.ClockInOnWork(ctx, in.User, time.Now())
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = v1.ErrorRecordExisted("error: clockin record existed")
		default:
			err = v1.ErrorUnknownError("error: unknown error")
		}
	}
	return &v1.ClockInOnWorkReply{}, err
}

func (svc *RecordService) ClockInOffWork(ctx context.Context, in *v1.ClockInOffWorkRequest) (*v1.ClockInOffWorkReply, error) {
	err := svc.uc.ClockInOffWork(ctx, in.User, time.Now())
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = v1.ErrorRecordExisted("error: clockin record existed")
		default:
			err = v1.ErrorUnknownError("error: unknown error")
		}
	}
	return &v1.ClockInOffWorkReply{}, err
}
