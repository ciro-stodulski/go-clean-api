package rabbitmq

import (
	"go-api/src/main/container"
	consumer "go-api/src/presentation/amqp/consumers"
)

type IAmqpServer interface {
	New(c *container.Container) IAmqpServer
	Start()
	StartConsumers(container []consumer.Comsumer, i int)
	NeedToReconnect(err error, msg string)
	LoadConsumers(c *container.Container) []consumer.Comsumer
}
