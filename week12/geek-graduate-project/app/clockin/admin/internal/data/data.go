package data

import (
	worktimev1 "clock-in/api/worktime/service/v1"
	"clock-in/app/clockin/admin/internal/conf"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWorkTimeRepo, NewWorkTimeClient)

// Data .
type Data struct {
	log *log.Helper
	wc  worktimev1.WorktimeServiceClient
}

func NewWorkTimeClient(config *conf.Data) worktimev1.WorktimeServiceClient {
	conn, err := grpc.Dial(
		config.WorktimeClient.Addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic("work time client initialize failed, reason: " + err.Error())
	}
	return worktimev1.NewWorktimeServiceClient(conn)
}

// NewData .
func NewData(
	c *conf.Data,
	wc worktimev1.WorktimeServiceClient,
	logger log.Logger,
) (*Data, error) {
	return &Data{
		log: log.NewHelper(log.With(logger, "module", "data")),
		wc:  wc,
	}, nil
}
