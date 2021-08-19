package comsumer

import (
	"go-api/src/main/container"
	comsumer "go-api/src/presentation/amqp/comsumers"
)

type findByIdConsumer struct {
	container container.Container
}

func New(c container.Container) comsumer.Comsumer {
	return &findByIdConsumer{c}
}
