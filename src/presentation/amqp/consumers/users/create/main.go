package create

import (
	"go-api/src/main/container"
	comsumer "go-api/src/presentation/amqp/consumers"
)

type createConsumer struct {
	container *container.Container
	queue     string
	schema    CreateDto
}

func NewConsumer(container *container.Container) comsumer.Comsumer {
	return &createConsumer{
		container: container,
		queue:     "create.user",
	}
}

func (createConsumer *createConsumer) GetQueue() string {
	return createConsumer.queue
}

func (createConsumer *createConsumer) GetSchema() interface{} {
	return createConsumer.schema
}
