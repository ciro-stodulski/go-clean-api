package amqpclient

import (
	rabbitmqadapter "go-clean-api/cmd/infra/adapters/rabbitmq"

	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/streadway/amqp"
)

type (
	ConfigAmqpClient struct {
		Exchange    string
		Routing_key string
	}

	AmqpClient interface {
		Publish(b []byte, c ConfigAmqpClient) error
	}
	amqpClient struct {
		channel *rabbitmq.Channel
	}
)

func New() AmqpClient {
	return &amqpClient{
		channel: rabbitmqadapter.GetChanel(),
	}
}

func (ampcc *amqpClient) Publish(body []byte, config ConfigAmqpClient) error {
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
