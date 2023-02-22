package verifynotificationusecase

import (
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-cases"
	"log"
)

type (
	notifyUseCase struct {
		NotificationService portsservice.NotificationService
	}
)

func New(ns portsservice.NotificationService) domainusecases.NotifyUserUseCase {
	return &notifyUseCase{
		NotificationService: ns,
	}
}

func (nuc *notifyUseCase) Notify(dto portsservice.Dto) error {

	log.Default().Println("amqp consumer completed with succeffully")

	nuc.NotificationService.CheckNotify(dto.Name)
	return nil
}
