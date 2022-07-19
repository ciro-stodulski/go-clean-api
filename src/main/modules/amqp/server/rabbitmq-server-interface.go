package rabbitmq

import (
	consumer "go-api/src/interface/amqp/consumers"
	"go-api/src/main/container"
)

type IAmqpServer interface {
	New(c *container.Container) IAmqpServer
	Start()
	StartConsumers(c []consumer.Comsumer, i int)
	NeedToReconnect(err error, msg string)
	LoadConsumers(c *container.Container) []consumer.Comsumer
}
