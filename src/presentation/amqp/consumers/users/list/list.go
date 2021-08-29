package list

import (
	ports_amqp "go-api/src/presentation/amqp/ports"
)

func (findByIdConsumer *listConsumer) MessageHandler(msg ports_amqp.Message) error {

	findByIdConsumer.container.ListUsersUseCase.ListUsers()
	return nil
}

func (findByIdConsumer *listConsumer) OnConsumerError(err error) error {
	return err
}
