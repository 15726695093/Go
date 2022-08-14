package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type RecordRepo interface {
	ClockInOnWork(context.Context) error
	ClockInOffWork(context.Context) error
}

type ClockInUsecase struct {
	record RecordRepo
	log    *log.Helper
}

func NewClockInUsecase(repo RecordRepo, logger log.Logger) *ClockInUsecase {
	return &ClockInUsecase{record: repo, log: log.NewHelper(logger)}
}

func (uc *ClockInUsecase) ClockInOnWork(ctx context.Context) error {
	return uc.record.ClockInOnWork(ctx)
}

func (uc *ClockInUsecase) ClockInOffWork(ctx context.Context) error {
	return uc.record.ClockInOffWork(ctx)
}
