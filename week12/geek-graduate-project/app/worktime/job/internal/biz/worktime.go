package biz

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

var ErrIncorrectFormat = errors.New("error: incorrect mesaage format")

type Record struct {
	User    int64
	Day1    int64
	Moment1 int64
	Day2    int64
	Moment2 int64
}

type WorkTimeRepo interface {
	CaculateWorktime(ctx context.Context, records *Record) error
}

type WorktimeUsecase struct {
	repo WorkTimeRepo
	log  *log.Helper
}

func NewWorktimeUsecase(repo WorkTimeRepo, logger log.Logger) *WorktimeUsecase {
	return &WorktimeUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *WorktimeUsecase) CaculateWorktime(ctx context.Context, message string) error {
	records, err := TransferMessage(message)
	if err != nil {
		return err
	}
	return uc.repo.CaculateWorktime(ctx, records)
}

func TransferMessage(message string) (*Record, error) {
	msg := strings.Split(message, "-")
	if len(msg) != 5 {
		return nil, ErrIncorrectFormat
	}
	user, err := strconv.Atoi(msg[0])
	if err != nil {
		return nil, ErrIncorrectFormat
	}
	day1, err := strconv.Atoi(msg[1])
	if err != nil {
		return nil, ErrIncorrectFormat
	}
	moment1, err := strconv.Atoi(msg[2])
	if err != nil {
		return nil, ErrIncorrectFormat
	}
	day2, err := strconv.Atoi(msg[3])
	if err != nil {
		return nil, ErrIncorrectFormat
	}
	moment2, err := strconv.Atoi(msg[4])
	if err != nil {
		return nil, ErrIncorrectFormat
	}
	return &Record{
		User:    int64(user),
		Day1:    int64(day1),
		Moment1: int64(moment1),
		Day2:    int64(day2),
		Moment2: int64(moment2),
	}, nil
}
