package container

import (
	deleteuserusecase "go-clean-api/cmd/application/use-case/delete-user"
	getuserusecase "go-clean-api/cmd/application/use-case/get-user"
	listusersusecase "go-clean-api/cmd/application/use-case/list-user"
	loadnewmessagingusecase "go-clean-api/cmd/application/use-case/load-new-messaging"
	registeruserusecase "go-clean-api/cmd/application/use-case/register-user"
	sendnewmessagingusecase "go-clean-api/cmd/application/use-case/send-new-messaging"
	verifynotificationusecase "go-clean-api/cmd/application/use-case/verify-notification"
	"go-clean-api/cmd/domain/dto"
	inputdto "go-clean-api/cmd/domain/dto/input"
	messagingentity "go-clean-api/cmd/domain/entity/messaging"
	"go-clean-api/cmd/domain/entity/user"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/main/container/factories"
)

type (
	Container struct {
		GetUserUseCase          usecase.UseCase[string, *user.User]
		RegisterUserUseCase     usecase.UseCase[dto.RegisterUser, *user.User]
		ListUsersUseCase        usecase.UseCase[any, any]
		DeleteUserUseCase       usecase.UseCase[string, any]
		NotifyUserUseCase       usecase.UseCase[dto.Event, any]
		SendNewMessagingUseCase usecase.UseCase[inputdto.MessagingInput, any]
		LoadNewMessagingUseCase usecase.UseCase[string, messagingentity.MessagingEntity]
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
		DeleteUserUseCase:       deleteuserusecase.New(userService),
		ListUsersUseCase:        listusersusecase.New(userService),
		NotifyUserUseCase:       verifynotificationusecase.New(notificationService),
		SendNewMessagingUseCase: sendnewmessagingusecase.New(containerConfig.SubjectIDChannels),
		LoadNewMessagingUseCase: loadnewmessagingusecase.New(containerConfig.SubjectIDChannels),
	}
}
