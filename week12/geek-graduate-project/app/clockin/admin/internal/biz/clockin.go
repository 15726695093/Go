package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
)

type WorkTime struct {
	Id     int64
	Day    int64
	Minute int64
}

type UserWorkTime struct {
	User     int64
	Worktime []*WorkTime
}

type WorkTimeRepo interface {
	GetUserWorkTime(ctx context.Context, user int64, day []int64) ([]*WorkTime, error)
}

type ClockinUsecase struct {
	repo WorkTimeRepo
	log  *log.Helper
}

func NewClockinUsecase(repo WorkTimeRepo, logger log.Logger) *ClockinUsecase {
	return &ClockinUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ClockinUsecase) GetWorkTime(ctx context.Context, user []int64, day []int64) ([]*UserWorkTime, error) {
	worktime := make([]*UserWorkTime, len(user))
	g, ctx := errgroup.WithContext(ctx)
	for k, v := range user {
		index := k
		userId := v
		g.Go(func() error {
			reply, err := uc.repo.GetUserWorkTime(ctx, userId, day)
			if err != nil {
				return err
			}
			worktime[index] = &UserWorkTime{
				User:     userId,
				Worktime: reply,
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return worktime, nil
}
