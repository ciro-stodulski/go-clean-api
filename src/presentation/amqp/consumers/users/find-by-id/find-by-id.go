package find_by_id

import (
	comsumer "go-api/src/presentation/amqp/consumers"

	"github.com/mitchellh/mapstructure"
)

func (findByIdConsumer *findByIdConsumer) MessageHandler(msg comsumer.Message) error {
	dto := findByIdConsumer.schema
	mapstructure.Decode(msg.Body, &dto)

	findByIdConsumer.container.ListUsersUseCase.ListUsers()
	return nil
}

func (findByIdConsumer *findByIdConsumer) OnConsumerError(err error) {

}
