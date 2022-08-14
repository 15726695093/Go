package biz

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Record struct {
	Id    int64
	User  int64
	Day   int64
	Type  int64
	Clock time.Time
}

type RecordRepo interface {
	ClockIn(ctx context.Context, record *Record) error
	CaculateWorkTime(ctx context.Context, user, day int64) error
}

type RecordUsecase struct {
	repo RecordRepo
	log  *log.Helper
}

func NewRecordUsecase(repo RecordRepo, logger log.Logger) *RecordUsecase {
	return &RecordUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RecordUsecase) ClockInOnWork(ctx context.Context, user int64, clock time.Time) error {
	day, err := getDate2Int(clock)
	if err != nil {
		return err
	}
	return uc.repo.ClockIn(ctx, &Record{
		User:  user,
		Day:   int64(day),
		Type:  1,
		Clock: clock,
	})
}

func (uc *RecordUsecase) ClockInOffWork(ctx context.Context, user int64, clock time.Time) error {
	day, err := getDate2Int(clock)
	if err != nil {
		return err
	}
	err = uc.repo.ClockIn(ctx, &Record{
		User:  user,
		Day:   int64(day),
		Type:  2,
		Clock: clock,
	})
	if err != nil {
		return err
	}
	// heat job for caculate work time
	return uc.repo.CaculateWorkTime(ctx, user, int64(day))
}

// record date part of timestamp into int64,
// for example time.now() => 20220101(int64)
func getDate2Int(clock time.Time) (int, error) {
	year, month, day := clock.Date()
	return strconv.Atoi(fmt.Sprintf("%d%02d%02d", year, month, day))
}
