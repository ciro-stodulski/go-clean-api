package verifyconsumer

import (
	"go-clean-api/cmd/domain/dto"
	usecase "go-clean-api/cmd/domain/use-case"
	consumer "go-clean-api/cmd/presentation/amqp/consumers"
	ports_amqp "go-clean-api/cmd/presentation/amqp/ports"
	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	Dto struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	}

	verifiyConsumer struct {
		NotifyUserUseCase usecase.UseCase[dto.Event, any]
	}
)

func New(NotifyUserUseCase usecase.UseCase[dto.Event, any]) consumer.Comsumer {
	return &verifiyConsumer{
		NotifyUserUseCase,
	}
}

func (createConsumer *verifiyConsumer) GetConfig() consumer.ConsumeConfig {
	return consumer.ConsumeConfig{
		Queue:  "notify.create",
		Schema: dto.Event{},
	}
}

func (createConsumer *verifiyConsumer) MessageHandler(msg ports_amqp.Message) error {
	dto := dto.Event{}

	mapstructure.Decode(msg.Body, &dto)

	_, err := createConsumer.NotifyUserUseCase.Perform(dto)

	return err
}

func (createConsumer *verifiyConsumer) OnConsumerError(err error) consumer.AckConfig {
	log.Default().Println("error verify consumer:", err)

	return consumer.AckConfig{
		Multiple: false,
		Requeue:  false,
	}
}
