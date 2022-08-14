package data

import (
	"clock-in/app/worktime/service/internal/conf"
	"clock-in/app/worktime/service/internal/data/ent"
	"clock-in/app/worktime/service/internal/data/ent/migrate"
	"context"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWorkTimeRepo)

// Data .
type Data struct {
	db *ent.Client
	rc *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, error) {
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		panic(err.Error())
	}
	if err := client.Schema.Create(context.TODO(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		panic(err.Error())
	}
	rc := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		DB:   0,
	})
	return &Data{
		db: client.Debug(),
		rc: rc,
	}, nil
}
