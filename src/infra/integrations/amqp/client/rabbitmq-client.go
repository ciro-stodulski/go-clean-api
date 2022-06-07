package amqp_client

import (
	types_client "go-api/src/infra/integrations/amqp/client/types"
	amqp_helper "go-api/src/main/module/amqp/rabbitmq/helper"
	"log"

	"github.com/streadway/amqp"
)

type AmqpClient struct {
	channel *amqp.Channel
}

func New() IAmqpClient {
	conn, err_connection := amqp.Dial(
		amqp_helper.GetConnection(),
	)

	failOnError(err_connection, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	return &AmqpClient{
		channel: ch,
	}
}

func (amqp_Client *AmqpClient) Publish(body []byte, config types_client.ConfigAmqpClient) error {
	err := amqp_Client.channel.Publish(
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
