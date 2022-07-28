package verifynotificationusecase

import (
	portsservice "go-api/cmd/core/ports"
	"log"
)

type (
	NotifyUserUseCase interface {
		Notify(dto portsservice.Dto) error
	}

	notifyUseCase struct {
		NotificationService portsservice.NotificationService
	}
)

func New(ns portsservice.NotificationService) NotifyUserUseCase {
	return &notifyUseCase{
		NotificationService: ns,
	}
}

func (nuc *notifyUseCase) Notify(dto portsservice.Dto) error {

	log.Default().Println("amqp consumer completed with succeffully")

	nuc.NotificationService.CheckNotify(dto.Name)

	return nil
}
