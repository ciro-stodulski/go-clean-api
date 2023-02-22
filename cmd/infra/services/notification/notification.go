package notificationservice

import (
	domainnotificationproducer "go-clean-api/cmd/domain/integrations/amqp"
	domainnotificationpbgrpc "go-clean-api/cmd/domain/integrations/grpc"
	domainnotificationcollection "go-clean-api/cmd/domain/repositories/no-sql"
	portsservice "go-clean-api/cmd/domain/services"
)

type (
	notificationService struct {
		NotificationProto      domainnotificationpbgrpc.NotificationPbGrpc
		NotificationProducer   domainnotificationproducer.NotificationProducer
		NotificationCollection domainnotificationcollection.NotificationCollection
	}
)

func New(pbs domainnotificationpbgrpc.NotificationPbGrpc, pn domainnotificationproducer.NotificationProducer, nc domainnotificationcollection.NotificationCollection) portsservice.NotificationService {
	return &notificationService{
		NotificationProto:      pbs,
		NotificationProducer:   pn,
		NotificationCollection: nc,
	}
}
