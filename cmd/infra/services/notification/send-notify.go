package notificationservice

import (
	notificationproducer "go-api/cmd/infra/integrations/amqp/notification"
)

func (ns notificationService) SendNotify(dto notificationproducer.Dto) error {
	return nil
}
