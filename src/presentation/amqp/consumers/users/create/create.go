package create

import (
	ports_amqp "go-api/src/presentation/amqp/ports"
	"log"

	"github.com/mitchellh/mapstructure"
)

func (createConsumer *createConsumer) MessageHandler(msg ports_amqp.Message) error {
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
