package find_by_id

import (
	"go-api/src/main/container"
	comsumer "go-api/src/presentation/amqp/comsumers"
)

type findByIdConsumer struct {
	container *container.Container
	queue     string
	schema    FindByIdDto
}

func NewConsumer(container *container.Container) comsumer.Comsumer {
	return &findByIdConsumer{
		container: container,
		queue:     "find.user",
	}
}

func (findByIdConsumer *findByIdConsumer) GetQueue() string {
	return findByIdConsumer.queue
}

func (findByIdConsumer *findByIdConsumer) GetSchema() interface{} {
	return findByIdConsumer.schema
}
