package amqpclient

import (
	"go-api/cmd/infra/adapters/rabbitmq"
	typesclient "go-api/cmd/infra/integrations/amqp/client/types"

	"github.com/streadway/amqp"
)

type amqpClient struct {
	channel *amqp.Channel
}

func New() AmqpClient {
	return &amqpClient{
		channel: rabbitmq.GetChanel(),
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
