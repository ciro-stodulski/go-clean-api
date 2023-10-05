package verifynotificationusecase

import (
	"go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/service"
	usecase "go-clean-api/cmd/domain/use-case"
	"log"
)

type (
	notifyUseCase struct {
		NotificationService service.NotificationService
	}
)

func New(ns service.NotificationService) usecase.UseCase[dto.Event, any] {
	return &notifyUseCase{
		NotificationService: ns,
	}
}

func (nuc *notifyUseCase) Perform(dto dto.Event) (any, error) {

	log.Default().Println("amqp consumer completed with succeffully")

	err := nuc.NotificationService.CheckNotify(dto.Name)

	return nil, err
}
