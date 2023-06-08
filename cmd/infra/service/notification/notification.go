package notificationservice

import (
	domainnotificationproducer "go-clean-api/cmd/domain/integration/amqp"
	domainnotificationpbgrpc "go-clean-api/cmd/domain/integration/grpc"
	domainnotificationcollection "go-clean-api/cmd/domain/repository/no-sql"
	service "go-clean-api/cmd/domain/service"
)

type (
	notificationService struct {
		NotificationProto      domainnotificationpbgrpc.NotificationPbGrpc
		NotificationProducer   domainnotificationproducer.NotificationProducer
		NotificationCollection domainnotificationcollection.NotificationCollection
	}
)

func New(pbs domainnotificationpbgrpc.NotificationPbGrpc, pn domainnotificationproducer.NotificationProducer, nc domainnotificationcollection.NotificationCollection) service.NotificationService {
	return &notificationService{
		NotificationProto:      pbs,
		NotificationProducer:   pn,
		NotificationCollection: nc,
	}
}
