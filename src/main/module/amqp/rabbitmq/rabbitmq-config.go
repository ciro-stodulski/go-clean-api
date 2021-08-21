package rabbitmq

import (
	"fmt"
	"go-api/src/main/container"
	consumer "go-api/src/presentation/amqp/comsumers"
	consumer_find_by_id "go-api/src/presentation/amqp/comsumers/users/find-by-id"
)

func (rabbit_mq *RabbitMq) LoadConsumers(container *container.Container) []consumer.Comsumer {
	fmt.Println(container)
	return []consumer.Comsumer{
		consumer_find_by_id.NewConsumer(container),
	}
}

func (rabbit_mq *RabbitMq) LoadProducers(c *container.Container) {

}
