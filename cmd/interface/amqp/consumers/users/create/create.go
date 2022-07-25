package usercreateconsumer

import (
	registeruserusecase "go-api/cmd/core/use-case/register-user"
	consumer "go-api/cmd/interface/amqp/consumers"
	ports_amqp "go-api/cmd/interface/amqp/ports"
	"go-api/cmd/main/container"
	"log"

	"github.com/mitchellh/mapstructure"
)

type createConsumer struct {
	container *container.Container
}

func New(c *container.Container) consumer.Comsumer {
	return &createConsumer{
		container: c,
	}
}

func (createConsumer *createConsumer) GetConfig() consumer.ConsumeConfig {
	return consumer.ConsumeConfig{
		Queue:  "create.user",
		Schema: registeruserusecase.Dto{},
	}
}

func (createConsumer *createConsumer) MessageHandler(msg ports_amqp.Message) error {
	dto := registeruserusecase.Dto{}
	mapstructure.Decode(msg.Body, dto)

	new_user, err := createConsumer.container.RegisterUserUseCase.Register(dto)

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
