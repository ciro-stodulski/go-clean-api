package notificationservice

import (
	portsservice "go-api/cmd/core/ports"
	"log"
)

func (ns notificationService) SendNotify(dto portsservice.Dto) error {

	ns.NotificationProducer.SendNotify(dto)

	log.Default().Println("Send notification with succeffully.")
	return nil
}
