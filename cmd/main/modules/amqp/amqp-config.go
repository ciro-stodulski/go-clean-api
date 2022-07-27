package amqp

import (
	consumer "go-api/cmd/interface/amqp/consumers"
	verifyconsumer "go-api/cmd/interface/amqp/consumers/notification/verify"
	"go-api/cmd/main/container"
)

func (rm *amqpModule) LoadConsumers(c *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		verifyconsumer.New(c),
	}
}
