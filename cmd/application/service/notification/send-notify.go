package notificationservice

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"log"
)

func (ns notificationService) SendNotify(dto domaindto.Event) error {

	ns.NotificationProducer.SendNotify(dto)

	log.Default().Println("Send notification with succeffully.")
	return nil
}
