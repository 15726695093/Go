package data

import (
	userv1 "clock-in/api/user/v1"
	"clock-in/app/clockin/interface/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.UserRepo = (*userRepo)(nil)

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	user, err := repo.data.uc.GetUserByName(ctx, &userv1.GetUserByNameRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return toBizUser(user.User), nil
}

func (repo *userRepo) Register(ctx context.Context, user *biz.User) (*biz.User, error) {
	newUser, err := repo.data.uc.SaveUser(ctx, &userv1.SaveUserRequest{
		User: toProtoUser(user),
	})
	if err != nil {
		return nil, err
	}
	return toBizUser(newUser.User), nil
}

func toBizUser(source *userv1.User) *biz.User {
	return &biz.User{
		Id:       source.Id,
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}

func toProtoUser(source *biz.User) *userv1.User {
	return &userv1.User{
		Id:       source.Id,
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}
