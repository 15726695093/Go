package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Phone    string
}

type UserRepo interface {
	GetUserById(context.Context, []int64) ([]*User, error)
	GetUserByName(context.Context, string) (*User, error)
	SearchUserByName(context.Context, string) ([]*User, error)
	SaveUser(context.Context, *User) (*User, error)
	RemoveUser(context.Context, []int64) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id []int64) ([]*User, error) {
	return uc.repo.GetUserById(ctx, id)
}

func (uc *UserUsecase) GetUserByName(ctx context.Context, name string) (*User, error) {
	return uc.repo.GetUserByName(ctx, name)
}

func (uc *UserUsecase) SearchUserByName(ctx context.Context, name string) ([]*User, error) {
	return uc.repo.SearchUserByName(ctx, name)
}

func (uc *UserUsecase) SaveUser(ctx context.Context, user *User) (*User, error) {
	return uc.repo.SaveUser(ctx, user)
}

func (uc *UserUsecase) RemoveUser(ctx context.Context, id []int64) error {
	return uc.repo.RemoveUser(ctx, id)
}
