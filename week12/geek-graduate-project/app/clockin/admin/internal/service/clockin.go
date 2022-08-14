package service

import (
	v1 "clock-in/api/clockin/admin/v1"
	"clock-in/app/clockin/admin/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService
type ClockinAdminService struct {
	v1.UnimplementedClockinAdminServiceServer

	uc  *biz.ClockinUsecase
	log *log.Helper
}

var _ v1.ClockinAdminServiceServer = (*ClockinAdminService)(nil)

// NewUserService new a greeter service.
func NewClockinServicee(uc *biz.ClockinUsecase, logger log.Logger) *ClockinAdminService {
	return &ClockinAdminService{uc: uc, log: log.NewHelper(logger)}
}

func (svc *ClockinAdminService) GetWorkTime(ctx context.Context, in *v1.GetWorkTimeRequest) (reply *v1.GetWorkTimeReply, err error) {
	worktime, err := svc.uc.GetWorkTime(ctx, in.User, in.Day)
	if err != nil {
		err = v1.ErrorUnknownError("unknown error")
	} else {
		reply = &v1.GetWorkTimeReply{
			Data: bulk2ProtoUserWorkTime(worktime),
		}
	}
	return
}

func bulk2ProtoUserWorkTime(source []*biz.UserWorkTime) []*v1.UserWorkTime {
	userWorkTime := make([]*v1.UserWorkTime, 0, len(source))
	for _, v := range source {
		userWorkTime = append(userWorkTime, &v1.UserWorkTime{
			User:     v.User,
			Worktime: bulk2ProtoWorkTime(v.Worktime),
		})
	}
	return userWorkTime
}

func bulk2ProtoWorkTime(source []*biz.WorkTime) []*v1.WorkTime {
	workTime := make([]*v1.WorkTime, 0, len(source))
	for _, v := range source {
		workTime = append(workTime, &v1.WorkTime{
			Day:    v.Day,
			Minute: v.Minute,
		})
	}
	return workTime
}
