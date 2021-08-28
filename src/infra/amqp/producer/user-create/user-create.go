package user_create

import (
	"go-api/src/infra/amqp/producer"
	types_client "go-api/src/main/module/amqp/rabbitmq/client/types"
)

func (userCreate *userCreate) Send(dto producer.Body) error {
	config := types_client.ConfigAmqpClient{
		Exchange:    userCreate.exchange,
		Routing_key: userCreate.routing_key,
	}

	err := userCreate.clientAmqp.Publish(
		[]byte(dto.Data.(string)),
		config,
	)

	return err
}
