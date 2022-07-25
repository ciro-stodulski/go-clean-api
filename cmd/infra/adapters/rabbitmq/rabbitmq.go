package rabbitmqadapter

import (
	"log"

	"github.com/streadway/amqp"
)

func GetChanel() *amqp.Channel {
	conn, err_conn := amqp.Dial(
		GetConnection(),
	)

	failOnError(err_conn, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	return ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
