package notificationservice

import (
	portsservice "go-api/cmd/core/ports"

	getuserservice "go-api/cmd/infra/integrations/grpc/user/get-user"
)

type (
	notificationService struct {
		NotificationService getuserservice.GetUserService
	}
)

func New(pbs getuserservice.GetUserService) portsservice.NotificationService {
	return &notificationService{
		NotificationService: pbs,
	}
}
