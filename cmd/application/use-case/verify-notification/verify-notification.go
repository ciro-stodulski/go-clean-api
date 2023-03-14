package verifynotificationusecase

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
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

func (nuc *notifyUseCase) Notify(dto domaindto.Event) (*domainexceptions.ApplicationException, error) {

	log.Default().Println("amqp consumer completed with succeffully")

	errApp, err := nuc.NotificationService.CheckNotify(dto.Name)

	return errApp, err
}
