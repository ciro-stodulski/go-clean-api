package notificationservice

import (
	portsservice "go-clean-api/cmd/core/ports"
	notificationproducer "go-clean-api/cmd/infra/integrations/amqp/notification"
	getuserservice "go-clean-api/cmd/infra/integrations/grpc/notification"
	notificationcollection "go-clean-api/cmd/infra/repositories/no-sql/notification"
)

type (
	notificationService struct {
		NotificationProto      getuserservice.NotificationPbGrpc
		NotificationProducer   notificationproducer.NotificationProducer
		NotificationCollection notificationcollection.NotificationCollection
	}
)

func New(pbs getuserservice.NotificationPbGrpc, pn notificationproducer.NotificationProducer, nc notificationcollection.NotificationCollection) portsservice.NotificationService {
	return &notificationService{
		NotificationProto:      pbs,
		NotificationProducer:   pn,
		NotificationCollection: nc,
	}
}
