package amqp

import (
	consumer "go-api/cmd/interface/amqp/consumers"
	consume_user_create "go-api/cmd/interface/amqp/consumers/users/create"
	"go-api/cmd/main/container"
)

func (rm *amqpModule) LoadConsumers(c *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		consume_user_create.New(c),
	}
}
