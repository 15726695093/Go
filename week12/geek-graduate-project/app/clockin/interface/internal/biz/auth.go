package biz

import (
	"context"
	"encoding/base64"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

var ErrIncorrectPassword = errors.New("error: incorrect password")

type User struct {
	Id       int64
	Name     string
	Password string
	Phone    string
}

type UserRepo interface {
	GetUserByName(context.Context, string) (*User, error)
	Register(context.Context, *User) (*User, error)
}

type AuthUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewAuthUsecase(repo UserRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) Login(ctx context.Context, name, password string) (string, error) {
	Auth, err := uc.repo.GetUserByName(ctx, name)
	if err != nil {
		return "", err
	}
	if password != Auth.Password {
		return "", ErrIncorrectPassword
	}
	return generateToken(Auth.Id, Auth.Name), nil
}

func (uc *AuthUsecase) Register(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Register(ctx, user)
}

func generateToken(id int64, name string) string {
	source := strconv.Itoa(int(id)) + "|" + name
	return base64.StdEncoding.EncodeToString([]byte(source))
}
