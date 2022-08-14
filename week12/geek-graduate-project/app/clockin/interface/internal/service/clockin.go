package service

import (
	v1 "clock-in/api/clockin/interface/v1"
	recordv1 "clock-in/api/record/v1"
	userv1 "clock-in/api/user/v1"
	"clock-in/app/clockin/interface/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService
type ClockInService struct {
	v1.UnimplementedClockinInterfaceServer
	au  *biz.AuthUsecase
	cu  *biz.ClockInUsecase
	log *log.Helper
}

var _ v1.ClockinInterfaceServer = (*ClockInService)(nil)

// NewUserService new a greeter service.
func NewClockInService(
	au *biz.AuthUsecase,
	cu *biz.ClockInUsecase,
	logger log.Logger,
) *ClockInService {
	return &ClockInService{
		au:  au,
		cu:  cu,
		log: log.NewHelper(logger),
	}
}

func (svc *ClockInService) Register(ctx context.Context, in *v1.RegisterRequest) (reply *v1.RegisterReply, err error) {
	user, err := svc.au.Register(ctx, &biz.User{
		Name:     in.Name,
		Password: in.Password,
		Phone:    in.Phone,
	})
	if err != nil {
		switch {
		case userv1.IsUserExist(err):
			err = v1.ErrorUserExistedError(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	} else {
		reply = &v1.RegisterReply{
			Id:    user.Id,
			Name:  user.Name,
			Phone: user.Phone,
		}
	}
	return
}

func (svc *ClockInService) Login(ctx context.Context, in *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	token, err := svc.au.Login(ctx, in.Username, in.Password)
	if err != nil {
		switch {
		case userv1.IsUserNotFound(err):
			err = v1.ErrorUserNotFound(err.Error())
		case err == biz.ErrIncorrectPassword:
			err = v1.ErrorPasswordIncorrect(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	} else {
		reply = &v1.LoginReply{
			Token: token,
		}
	}
	return
}

func (svc *ClockInService) ClockinOnWork(ctx context.Context, in *v1.ClockinOnWorkRequest) (*v1.ClockinOnWorkReply, error) {
	err := svc.cu.ClockInOnWork(ctx)
	if err != nil {
		switch {
		case recordv1.IsRecordExisted(err):
			err = v1.ErrorRecordExisted(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.ClockinOnWorkReply{}, err
}

func (svc *ClockInService) ClockinOffWork(ctx context.Context, in *v1.ClockinOffWorkRequest) (*v1.ClockinOffWorkReply, error) {
	err := svc.cu.ClockInOffWork(ctx)
	if err != nil {
		switch {
		case recordv1.IsRecordExisted(err):
			err = v1.ErrorRecordExisted(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	return &v1.ClockinOffWorkReply{}, err
}
