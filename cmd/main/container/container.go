package container

import (
	delete_user "go-clean-api/cmd/core/use-case/delete-user"
	get_user_use_case "go-clean-api/cmd/core/use-case/get-user"
	list_users "go-clean-api/cmd/core/use-case/list-user"
	registeruserusecase "go-clean-api/cmd/core/use-case/register-user"
	verifynotificationusecase "go-clean-api/cmd/core/use-case/verify-notification"
	"go-clean-api/cmd/main/container/factories"
)

type (
	Container struct {
		GetUserUseCase      get_user_use_case.GetUserUseCase
		RegisterUserUseCase registeruserusecase.RegisterUserUseCase
		ListUsersUseCase    list_users.ListUsersUseCase
		DeleteUserUseCase   delete_user.DeleteUserUseCase
		VerifyUseCase       verifynotificationusecase.NotifyUserUseCase
	}
)

func New() *Container {
	container_config := newContainerConfig()

	infra_context := factories.MakeInfraContext(
		container_config.Grpc_client,
		container_config.Amqp_client,
		container_config.Http_client,
		container_config.Database,
		container_config.Cache_client)

	user_service := factories.MakeServiceContext(infra_context).User_service

	notification_service := factories.MakeServiceContext(infra_context).Notification_service

	return &Container{
		GetUserUseCase: get_user_use_case.New(
			user_service,
		),
		RegisterUserUseCase: registeruserusecase.New(
			user_service, notification_service,
		),
		DeleteUserUseCase: delete_user.New(user_service),
		ListUsersUseCase:  list_users.New(user_service),
		VerifyUseCase:     verifynotificationusecase.New(notification_service),
	}
}
