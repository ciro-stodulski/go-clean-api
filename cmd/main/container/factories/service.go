package factories

import (
	portsservice "go-clean-api/cmd/core/ports"
	notificationService "go-clean-api/cmd/infra/services/notification"
	userservice "go-clean-api/cmd/infra/services/user"
)

type (
	ServiceCaseContext struct {
		User_service         portsservice.UserService
		Notification_service portsservice.NotificationService
	}
)

func MakeServiceContext(infra_context InfraContext) ServiceCaseContext {
	return ServiceCaseContext{
		User_service: userservice.New(
			infra_context.User_repository,
			infra_context.Json_place_holder_integration,
			infra_context.Users_cache,
		),
		Notification_service: notificationService.New(
			infra_context.NotificationPbGrpc,
			infra_context.Notification_amqp,
			infra_context.Notification_collection,
		),
	}
}
