package data

import (
	"clock-in/app/record/service/internal/conf"
	"clock-in/app/record/service/internal/data/ent"
	"clock-in/app/record/service/internal/data/ent/migrate"
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRecordRepo)

// Data .
type Data struct {
	db *ent.Client
	mq *amqp.Channel
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
		mq: NewMQPublisher(c),
	}, nil
}

func NewMQPublisher(c *conf.Data) *amqp.Channel {
	conn, err := amqp.Dial(c.Rabbitmq.Addr)
	if err != nil {
		panic(err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	_, err = ch.QueueDeclare(
		"worktime", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		panic(err.Error())
	}
	return ch
}
