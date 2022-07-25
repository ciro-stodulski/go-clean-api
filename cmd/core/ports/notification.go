package portsservice

import (
	notificationproducer "go-api/cmd/infra/integrations/amqp/producer/notification"
)

type (
	NotificationService interface {
		SendNotify(dto notificationproducer.Dto) error
		CheckNotify(msg string) (string error)
	}
)
