package create

import (
	registeruserusecase "go-api/cmd/core/use-case/register-user"
	comsumer "go-api/cmd/interface/amqp/consumers"
	"go-api/cmd/main/container"
)

type createConsumer struct {
	container *container.Container
	queue     string
	schema    registeruserusecase.Dto
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
