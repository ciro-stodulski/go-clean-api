package create

import (
	comsumer "go-api/src/presentation/amqp/consumers"

	"github.com/mitchellh/mapstructure"
)

func (createConsumer *createConsumer) MessageHandler(msg comsumer.Message) error {
	dto := createConsumer.schema
	mapstructure.Decode(msg.Body, &dto)

	createConsumer.container.ListUsersUseCase.ListUsers()
	return nil
}

func (createConsumer *createConsumer) OnConsumerError(err error) {

}
