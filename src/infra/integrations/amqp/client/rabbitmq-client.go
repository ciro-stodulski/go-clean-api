package amqpclient

import (
	typesclient "go-api/src/infra/integrations/amqp/client/types"
	amqphelper "go-api/src/main/module/amqp/helper"
	"log"

	"github.com/streadway/amqp"
)

type amqpClient struct {
	channel *amqp.Channel
}

func New() AmqpClient {
	conn, err_conn := amqp.Dial(
		amqphelper.GetConnection(),
	)

	failOnError(err_conn, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	return &amqpClient{
		channel: ch,
	}
}

func (ampcc *amqpClient) Publish(body []byte, config typesclient.ConfigAmqpClient) error {
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
