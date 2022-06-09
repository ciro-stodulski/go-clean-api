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

func NewConsumer(c *container.Container) comsumer.Comsumer {
	return &listConsumer{
		container: c,
		queue:     "list.user",
	}
}

func (lc *listConsumer) GetQueue() string {
	return lc.queue
}

func (lc *listConsumer) GetSchema() interface{} {
	return lc.schema
}
