package amqp

import (
	"go-api/src/main/container"
	types_client "go-api/src/main/module/amqp/rabbitmq/client/types"
	consumer "go-api/src/presentation/amqp/consumers"
)

type AmqpServer interface {
	New(c *container.Container) AmqpServer
	Start()
	StartConsumers(container []consumer.Comsumer, i int)
	NeedToReconnect(err error, msg string)
	LoadConsumers(c *container.Container) []consumer.Comsumer
}

type AmqpClient interface {
	New(config types_client.ConfigAmqpClient) AmqpClient
	Publish(body []byte) error
}
