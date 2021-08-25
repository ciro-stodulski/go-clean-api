package find_by_id

import (
	ports_amqp "go-api/src/presentation/amqp/ports"

	"github.com/mitchellh/mapstructure"
)

func (findByIdConsumer *findByIdConsumer) MessageHandler(msg ports_amqp.Message) error {
	dto := findByIdConsumer.schema
	mapstructure.Decode(msg.Body, &dto)

	findByIdConsumer.container.ListUsersUseCase.ListUsers()
	return nil
}

func (findByIdConsumer *findByIdConsumer) OnConsumerError(err error) error {
	return err
}
