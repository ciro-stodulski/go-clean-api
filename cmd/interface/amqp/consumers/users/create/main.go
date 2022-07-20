package create

import (
	comsumer "go-api/cmd/interface/amqp/consumers"
	create_dto "go-api/cmd/interface/amqp/consumers/users/create/dto"
	"go-api/cmd/main/container"
)

type createConsumer struct {
	container *container.Container
	queue     string
	schema    create_dto.CreateDto
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
