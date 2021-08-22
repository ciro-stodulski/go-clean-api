package rabbitmq

import (
	"go-api/src/main/container"
	consumer "go-api/src/presentation/amqp/consumers"
	consume_user_create "go-api/src/presentation/amqp/consumers/users/create"
	consumer_user_find_by_id "go-api/src/presentation/amqp/consumers/users/find-by-id"
)

func (rabbit_mq *RabbitMq) LoadConsumers(container *container.Container) []consumer.Comsumer {
	return []consumer.Comsumer{
		consume_user_create.NewConsumer(container),
		consumer_user_find_by_id.NewConsumer(container),
	}
}

func (rabbit_mq *RabbitMq) LoadProducers(c *container.Container) {

}
