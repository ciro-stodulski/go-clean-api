package container

import (
	delete_user "go-clean-api/cmd/application/use-case/delete-user"
	get_user_use_case "go-clean-api/cmd/application/use-case/get-user"
	list_users "go-clean-api/cmd/application/use-case/list-user"
	registeruserusecase "go-clean-api/cmd/application/use-case/register-user"
	verifynotificationusecase "go-clean-api/cmd/application/use-case/verify-notification"
	"go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/main/container/factories"
)

type (
	Container struct {
		GetUserUseCase      usecase.IUseCase[string, *user.User]
		RegisterUserUseCase usecase.IUseCase[dto.RegisterUser, *user.User]
		ListUsersUseCase    usecase.IUseCase[interface{}, interface{}]
		DeleteUserUseCase   usecase.IUseCase[string, interface{}]
		NotifyUserUseCase   usecase.IUseCase[dto.Event, interface{}]
	}
)

func New() *Container {
	container_config := newContainerConfig()

	infra_context := factories.MakeInfraContext(
		container_config.Grpc_client,
		container_config.Amqp_client,
		container_config.Http_client,
		container_config.Database,
		container_config.Cache_client,
		container_config.DatabaseNoSql)

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
		NotifyUserUseCase: verifynotificationusecase.New(notification_service),
	}
}
