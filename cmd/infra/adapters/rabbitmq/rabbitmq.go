package rabbitmqadapter

import (
	"log"

	"github.com/isayme/go-amqp-reconnect/rabbitmq"
)

func GetChanel() *rabbitmq.Channel {
	rabbitmq.Debug = true

	conn, err_conn := rabbitmq.Dial(
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
