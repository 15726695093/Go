package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

type WorkTime struct {
	Id     int64
	Day    int64
	Minute int64
}

type Record struct {
	Day    int64
	User   int64
	Moment int64
}

var ErrMust2Record = errors.New("error: must send two records")
var ErrMustSameDay = errors.New("error: records must belongs to same day")

type WorkTimeRepo interface {
	GetUserWorkTime(ctx context.Context, user int64, day []int64) ([]*WorkTime, error)
	CreateWorkTime(ctx context.Context, user int64, records []*Record) error
}

type WorkTimeUsecase struct {
	repo WorkTimeRepo
	log  *log.Helper
}

func NewWorkTimeUsecase(repo WorkTimeRepo, logger log.Logger) *WorkTimeUsecase {
	return &WorkTimeUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *WorkTimeUsecase) GetUserWorkTime(ctx context.Context, user int64, day []int64) ([]*WorkTime, error) {
	return uc.repo.GetUserWorkTime(ctx, user, day)
}

func (uc *WorkTimeUsecase) CreateWorkTime(ctx context.Context, user int64, records []*Record) error {
	if len(records) != 2 {
		return ErrMust2Record
	}
	if records[0].Day != records[1].Day {
		return ErrMustSameDay
	}
	return uc.repo.CreateWorkTime(ctx, user, records)
}
