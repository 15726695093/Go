package data

import (
	"clock-in/app/worktime/job/internal/conf"

	worktimev1 "clock-in/api/worktime/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWorkTimeClient, NewWorkTimeRepo)

// Data .
type Data struct {
	log *log.Helper
	wc  worktimev1.WorktimeServiceClient
}

// NewData .
func NewData(c *conf.Data, wc worktimev1.WorktimeServiceClient, logger log.Logger) (*Data, error) {
	return &Data{
		wc:  wc,
		log: log.NewHelper(log.With(logger, "module", "data")),
	}, nil
}

func NewWorkTimeClient(config *conf.Data) worktimev1.WorktimeServiceClient {
	conn, err := grpc.Dial(
		config.WorkTimeClient.Addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err.Error())
	}
	return worktimev1.NewWorktimeServiceClient(conn)
}
