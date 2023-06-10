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

func New(ns service.NotificationService) usecase.NotifyUserUseCase {
	return &notifyUseCase{
		NotificationService: ns,
	}
}

func (nuc *notifyUseCase) Notify(dto dto.Event) error {

	log.Default().Println("amqp consumer completed with succeffully")

	err := nuc.NotificationService.CheckNotify(dto.Name)

	return err
}
