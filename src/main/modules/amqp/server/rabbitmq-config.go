package rabbitmq

import (
	consumer "go-api/src/interface/amqp/consumers"
	consume_user_create "go-api/src/interface/amqp/consumers/users/create"
	consumer_user_list "go-api/src/interface/amqp/consumers/users/list"
	"go-api/src/main/container"
)

func (rm *RabbitMq) LoadConsumers(c *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		consume_user_create.NewConsumer(c),
		consumer_user_list.NewConsumer(c),
	}
}
