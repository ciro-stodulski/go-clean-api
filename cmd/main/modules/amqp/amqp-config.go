package amqp

import (
	consumer "go-clean-api/cmd/interface/amqp/consumers"
	verifyconsumer "go-clean-api/cmd/interface/amqp/consumers/notification/verify"
	"go-clean-api/cmd/main/container"
)

func (rm *amqpModule) LoadConsumers(c *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		verifyconsumer.New(c),
	}
}
