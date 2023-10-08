package factories

import (
	notificationService "go-clean-api/cmd/application/service/notification"
	userservice "go-clean-api/cmd/application/service/user"
	service "go-clean-api/cmd/domain/service"
)

type (
	ServiceCaseContext struct {
		UserService         service.UserService
		NotificationService service.NotificationService
	}
)

func MakeServiceContext(infraContext InfraContext) ServiceCaseContext {
	return ServiceCaseContext{
		UserService: userservice.New(
			infraContext.User_repository,
			infraContext.Json_place_holder_integration,
			infraContext.Users_cache,
		),
		NotificationService: notificationService.New(
			infraContext.NotificationPbGrpc,
			infraContext.Notification_amqp,
			infraContext.Notification_collection,
		),
	}
}
