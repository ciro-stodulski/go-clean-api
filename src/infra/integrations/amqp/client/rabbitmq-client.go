package amqpclient

import (
	typesclient "go-api/src/infra/integrations/amqp/client/types"
	amqphelper "go-api/src/main/module/amqp/rabbitmq/helper"
	"log"

	"github.com/streadway/amqp"
)

type AmqpClient struct {
	channel *amqp.Channel
}

func New() IAmqpClient {
	conn, err_connection := amqp.Dial(
		amqphelper.GetConnection(),
	)

	failOnError(err_connection, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	return &AmqpClient{
		channel: ch,
	}
}

func (ampcc *AmqpClient) Publish(body []byte, config typesclient.ConfigAmqpClient) error {
	err := ampcc.channel.Publish(
		config.Exchange,    // exchange
		config.Routing_key, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)

	return err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
