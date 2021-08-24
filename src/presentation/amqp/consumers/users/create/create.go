package create

import (
	comsumer "go-api/src/presentation/amqp/consumers"
	"log"

	"github.com/mitchellh/mapstructure"
)

func (createConsumer *createConsumer) MessageHandler(msg comsumer.Message) error {
	dto := createConsumer.schema
	mapstructure.Decode(msg.Body, &dto)

	new_user, err := createConsumer.container.CreateUserUseCase.CreateUser(dto)

	if err != nil {
		return err
	}

	log.Default().Println("User register with succeffully: user-", new_user.Name)

	return nil
}

func (createConsumer *createConsumer) OnConsumerError(err error) error {
	log.Default().Println("error:", err)

	return err
}
