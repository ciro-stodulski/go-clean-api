package verifyconsumer

import (
	portsservice "go-clean-api/cmd/core/ports"
	consumer "go-clean-api/cmd/interface/amqp/consumers"
	ports_amqp "go-clean-api/cmd/interface/amqp/ports"
	"go-clean-api/cmd/main/container"
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

func (createConsumer *verifiyConsumer) OnConsumerError(err error) consumer.AckConfig {
	log.Default().Println("error verify consumer:", err)

	return consumer.AckConfig{
		Multiple: false,
		Requeue:  false,
	}
}
