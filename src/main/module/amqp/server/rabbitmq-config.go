package rabbitmq

import (
	"go-api/src/main/container"
	consumer "go-api/src/presentation/amqp/consumers"
	consume_user_create "go-api/src/presentation/amqp/consumers/users/create"
	consumer_user_list "go-api/src/presentation/amqp/consumers/users/list"
)

func (rm *RabbitMq) LoadConsumers(c *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		consume_user_create.NewConsumer(c),
		consumer_user_list.NewConsumer(c),
	}
}
