package factories

import (
	service "go-clean-api/cmd/domain/service"
	notificationService "go-clean-api/cmd/infra/service/notification"
	userservice "go-clean-api/cmd/infra/service/user"
)

type (
	ServiceCaseContext struct {
		User_service         service.UserService
		Notification_service service.NotificationService
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
