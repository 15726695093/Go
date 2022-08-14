package service

import (
	"context"

	v1 "clock-in/api/user/v1"
	"clock-in/app/user/service/internal/biz"
	"clock-in/app/user/service/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService
type UserService struct {
	v1.UnimplementedUserServiceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

var _ v1.UserServiceServer = (*UserService)(nil)

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) GetUserById(ctx context.Context, in *v1.GetUserByIdRequest) (reply *v1.GetUserByIdReply, err error) {
	users, err := s.uc.GetUserById(ctx, in.Id)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			err = v1.ErrorUserNotFound("error: get user failed, reason: user not found")
		default:
			err = v1.ErrorUnknownError("error: get user failed, reason: unknown error")
		}
	} else {
		reply = &v1.GetUserByIdReply{
			User: bulk2ProtoUser(users),
		}
	}
	return
}

func (s *UserService) GetUserByName(ctx context.Context, in *v1.GetUserByNameRequest) (reply *v1.GetUserByNameReply, err error) {
	user, err := s.uc.GetUserByName(ctx, in.Name)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			err = v1.ErrorUserNotFound("error: get user failed, reason: user not found")
		default:
			err = v1.ErrorUnknownError("error: get user failed, reason: unknown error")
		}
	} else {
		reply = &v1.GetUserByNameReply{
			User: toProtoUser(user),
		}
	}
	return
}

func (s *UserService) SearchUserByName(ctx context.Context, in *v1.SearchUserByNameRequest) (reply *v1.SearchUserByNameReply, err error) {
	users, err := s.uc.SearchUserByName(ctx, in.Name)
	if err != nil {
		err = v1.ErrorUnknownError("error: search user failed, reason: unknown error")
	} else {
		reply = &v1.SearchUserByNameReply{
			User: bulk2ProtoUser(users),
		}
	}
	return
}

func (s *UserService) SaveUser(ctx context.Context, in *v1.SaveUserRequest) (reply *v1.SaveUserReply, err error) {
	user, err := s.uc.SaveUser(ctx, toBizUser(in.User))
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = v1.ErrorUserExist("error: save user failed, reason: user name existed")
		default:
			err = v1.ErrorUnknownError("error: save user failed, reason: unknown error")
		}
	} else {
		reply = &v1.SaveUserReply{
			User: toProtoUser(user),
		}
	}
	return reply, err
}

func (s *UserService) RemoveUser(ctx context.Context, in *v1.RemoveUserRequest) (reply *v1.RemoveUserReply, err error) {
	err = s.uc.RemoveUser(ctx, in.Id)
	if err != nil {
		err = v1.ErrorUnknownError("error: remove user failed, reason: unknown error")
	}
	return &v1.RemoveUserReply{}, err
}

func toProtoUser(source *biz.User) *v1.User {
	return &v1.User{
		Id:       int64(source.Id),
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}

func toBizUser(source *v1.User) *biz.User {
	return &biz.User{
		Id:       source.Id,
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}

func bulk2ProtoUser(source []*biz.User) []*v1.User {
	user := make([]*v1.User, 0, len(source))
	for _, v := range source {
		user = append(user, toProtoUser(v))
	}
	return user
}
