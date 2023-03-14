package notificationservice

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"log"
)

func (ns notificationService) SendNotify(dto domaindto.Event) (*domainexceptions.ApplicationException, error) {

	ns.NotificationProducer.SendNotify(dto)

	log.Default().Println("Send notification with succeffully.")
	return nil, nil
}
