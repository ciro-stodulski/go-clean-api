package verifyconsumer

import (
	portsservice "go-api/cmd/core/ports"
	consumer "go-api/cmd/interface/amqp/consumers"
	ports_amqp "go-api/cmd/interface/amqp/ports"
	"go-api/cmd/main/container"
	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	Dto struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	}

	verifiyConsumer struct {
		container *container.Container
	}
)

func New(c *container.Container) consumer.Comsumer {
	return &verifiyConsumer{
		container: c,
	}
}

func (createConsumer *verifiyConsumer) GetConfig() consumer.ConsumeConfig {
	return consumer.ConsumeConfig{
		Queue:  "notify.create",
		Schema: portsservice.Dto{},
	}
}

func (createConsumer *verifiyConsumer) MessageHandler(msg ports_amqp.Message) error {
	dto := portsservice.Dto{}

	mapstructure.Decode(msg.Body, &dto)

	err := createConsumer.container.VerifyUseCase.Notify(dto)

	if err != nil {
		return err
	}

	return nil
}

func (createConsumer *verifiyConsumer) OnConsumerError(err error) error {
	log.Default().Println("error:", err)

	return err
}
