package list

import (
	ports_amqp "go-api/src/presentation/amqp/ports"
)

func (lc *listConsumer) MessageHandler(msg ports_amqp.Message) (err error) {

	lc.container.ListUsersUseCase.ListUsers()
	return
}

func (lc *listConsumer) OnConsumerError(err error) error {
	return err
}
