package amqp

import (
	"go-api/src/main/container"
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
	New() AmqpClient
}
