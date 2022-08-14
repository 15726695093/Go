package data

import (
	"clock-in/app/user/service/internal/conf"
	"clock-in/app/user/service/internal/data/ent"
	"clock-in/app/user/service/internal/data/ent/migrate"
	"context"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
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
	return &Data{
		db: client.Debug(),
	}, nil
}
