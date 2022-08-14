package job

import (
	"testing"

	"github.com/streadway/amqp"
)

func TestCaculateWorktime(t *testing.T) {
	conn, err := amqp.Dial("amqp://root:000000@127.0.0.1:5672/")
	if err != nil {
		t.Error(err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Error(err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"worktime", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		t.Error(err.Error())
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("5-20220119-1642644943-20220119-1642673743"),
		})
	if err != nil {
		t.Error(err.Error())
	}
}
