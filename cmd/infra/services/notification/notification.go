package notificationservice

import (
	portsservice "go-clean-api/cmd/core/ports"

	notificationproducer "go-clean-api/cmd/infra/integrations/amqp/notification"
	getuserservice "go-clean-api/cmd/infra/integrations/grpc/notification"
)

type (
	notificationService struct {
		NotificationProto    getuserservice.NotificationPbGrpc
		NotificationProducer notificationproducer.NotificationProducer
	}
)

func New(pbs getuserservice.NotificationPbGrpc, pn notificationproducer.NotificationProducer) portsservice.NotificationService {
	return &notificationService{
		NotificationProto:    pbs,
		NotificationProducer: pn,
	}
}
