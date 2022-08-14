package mq

import (
	"clock-in/app/worktime/job/internal/conf"

	"github.com/streadway/amqp"
)

func NewConsumer(c *conf.Data) <-chan amqp.Delivery {
	conn, err := amqp.Dial(c.Rabbitmq.Addr)
	if err != nil {
		panic("error: rabbit mq connect failed, reason: " + err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		panic("error: rabbit mq channel connect failed, reason: " + err.Error())
	}
	q, err := ch.QueueDeclare(
		"worktime", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		panic("error: rabbit mq declare queue failed, reason: " + err.Error())
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic("error: rabbit mq consume connect failed, reason: " + err.Error())
	}
	return msgs
}
