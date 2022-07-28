package container

import (
	delete_user "go-api/cmd/core/use-case/delete-user"
	get_user_use_case "go-api/cmd/core/use-case/get-user"
	list_users "go-api/cmd/core/use-case/list-user"
	registeruserusecase "go-api/cmd/core/use-case/register-user"
	verifynotificationusecase "go-api/cmd/core/use-case/verify-notification"
	amqpclient "go-api/cmd/infra/integrations/amqp"
	notificationproducer "go-api/cmd/infra/integrations/amqp/notification"
	json_place_holder "go-api/cmd/infra/integrations/http/jsonplaceholder"
	usersjsonplaceholder "go-api/cmd/infra/repositories/cache/users-jsonplaceholder"
	model_user "go-api/cmd/infra/repositories/sql/user"
	"go-api/cmd/shared/env"

	grpc_client "go-api/cmd/infra/integrations/grpc"
	notificationpbgrpc "go-api/cmd/infra/integrations/grpc/notification"
	"go-api/cmd/infra/integrations/grpc/notification/pb"
	http_service "go-api/cmd/infra/integrations/http"
	cache_client "go-api/cmd/infra/repositories/cache"
	notificationService "go-api/cmd/infra/services/notification"
	userservice "go-api/cmd/infra/services/user"

	database "go-api/cmd/infra/adapters/mysql"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		GetUserUseCase      get_user_use_case.GetUserUseCase
		RegisterUserUseCase registeruserusecase.RegisterUserUseCase
		ListUsersUseCase    list_users.ListUsersUseCase
		DeleteUserUseCase   delete_user.DeleteUserUseCase
		VerifyUseCase       verifynotificationusecase.NotifyUserUseCase
	}
)

var db database.MysqlAdapter

func New() *Container {
	db.ConnectToDatabase()

	grpc_client := grpc_client.New()
	find_user_service := notificationpbgrpc.New(
		pb.NewNotificationPbClient(
			grpc_client.GetConnection(
				env.Env().GrpcClientUrl)))

	amqp_client := amqpclient.New()
	notification_amqp := notificationproducer.New(amqp_client)

	http_service := http_service.New()
	json_place_holder_integration := json_place_holder.New(http_service)

	user_repository := model_user.NewUserRepository(db.Db)

	cache_client := cache_client.New()
	users_cache := usersjsonplaceholder.New(cache_client)

	user_service := userservice.New(user_repository, json_place_holder_integration, users_cache)
	notification_service := notificationService.New(find_user_service, notification_amqp)

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
