package amqp_client

import (
	amqp_client "go-api/src/main/module/amqp/rabbitmq"
	types_client "go-api/src/main/module/amqp/rabbitmq/client/types"
	"log"

	"github.com/streadway/amqp"
)

type AmqpClient struct {
	channel     *amqp.Channel
	exchange    string
	routing_key string
}

func (amqp_Client *AmqpClient) New(config types_client.ConfigAmqpClient) amqp_client.AmqpClient {
	conn, err_connection := amqp.Dial(
		amqp_client.GetConnection(),
	)

	failOnError(err_connection, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	return &AmqpClient{
		channel:     ch,
		exchange:    config.Exchange,
		routing_key: config.Routing_key,
	}
}

func (amqp_Client *AmqpClient) Publish(body []byte) error {
	err := amqp_Client.channel.Publish(
		amqp_Client.exchange,    // exchange
		amqp_Client.routing_key, // routing key
		false,                   // mandatory
		false,                   // immediate
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
