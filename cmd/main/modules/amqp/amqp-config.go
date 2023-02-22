package amqp

import (
	"go-clean-api/cmd/main/container"
	consumer "go-clean-api/cmd/presetation/amqp/consumers"
	verifyconsumer "go-clean-api/cmd/presetation/amqp/consumers/notification/verify"
)

func (rm *amqpModule) LoadConsumers(c *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		verifyconsumer.New(c),
	}
}
