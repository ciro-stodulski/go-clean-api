package container

import (
	deleteuserusecase "go-clean-api/cmd/application/use-case/delete-user"
	getuserusecase "go-clean-api/cmd/application/use-case/get-user"
	listusersusecase "go-clean-api/cmd/application/use-case/list-user"
	registeruserusecase "go-clean-api/cmd/application/use-case/register-user"
	verifynotificationusecase "go-clean-api/cmd/application/use-case/verify-notification"
	"go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/main/container/factories"
)

type (
	Container struct {
		GetUserUseCase      usecase.UseCase[string, *user.User]
		RegisterUserUseCase usecase.UseCase[dto.RegisterUser, *user.User]
		ListUsersUseCase    usecase.UseCase[interface{}, interface{}]
		DeleteUserUseCase   usecase.UseCase[string, interface{}]
		NotifyUserUseCase   usecase.UseCase[dto.Event, interface{}]
	}
)

func New() *Container {
	containerConfig := newContainerConfig()

	infraContext := factories.MakeInfraContext(
		containerConfig.GrpcClient,
		containerConfig.AmqpClient,
		containerConfig.HttpClient,
		containerConfig.Database,
		containerConfig.CacheClient,
		containerConfig.DatabaseNoSql)

	userService := factories.MakeServiceContext(infraContext).UserService

	notificationService := factories.MakeServiceContext(infraContext).NotificationService

	return &Container{
		GetUserUseCase: getuserusecase.New(
			userService,
		),
		RegisterUserUseCase: registeruserusecase.New(
			userService, notificationService,
		),
		DeleteUserUseCase: deleteuserusecase.New(userService),
		ListUsersUseCase:  listusersusecase.New(userService),
		NotifyUserUseCase: verifynotificationusecase.New(notificationService),
	}
}
