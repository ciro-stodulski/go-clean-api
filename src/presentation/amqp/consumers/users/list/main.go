package list

import (
	"go-api/src/main/container"
	comsumer "go-api/src/presentation/amqp/consumers"
)

type listConsumer struct {
	container *container.Container
	queue     string
	schema    interface{}
}

func NewConsumer(container *container.Container) comsumer.Comsumer {
	return &listConsumer{
		container: container,
		queue:     "list.user",
	}
}

func (findByIdConsumer *listConsumer) GetQueue() string {
	return findByIdConsumer.queue
}

func (findByIdConsumer *listConsumer) GetSchema() interface{} {
	return findByIdConsumer.schema
}
