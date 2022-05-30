package container

import (
	create_user_use_case "go-api/src/core/useCases/create-user"
	create_user_producer_use_case "go-api/src/core/useCases/create-user-producer"
	delete_user "go-api/src/core/useCases/delete-user"
	get_user_use_case "go-api/src/core/useCases/get-user"
	get_user_grpc "go-api/src/core/useCases/get-user-grpc"
	list_users "go-api/src/core/useCases/list-user"
	"os"

	users_cache "go-api/src/infra/cache/users"
	create_user_amqp "go-api/src/infra/integrations/amqp/producer/user-create"
	json_place_holder "go-api/src/infra/integrations/http/jsonplaceholder"
	model_user "go-api/src/infra/repositories/user"

	grpc_client "go-api/src/infra/integrations/grpc/client"
	find_user_service "go-api/src/infra/integrations/grpc/user/get-user"
	"go-api/src/infra/integrations/grpc/user/get-user/pb"
	amqp_client "go-api/src/main/module/amqp/rabbitmq/client"
	cache_client "go-api/src/main/module/cache/redis"
	http_service "go-api/src/main/module/http/client"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		GetUserUseCase            get_user_use_case.GetUserUseCase
		GetUserGrpcUseCase        get_user_grpc.GetUserGrpcUseCase
		CreateUserUseCase         create_user_use_case.CreateUserUseCase
		CreateUserProducerUseCase create_user_producer_use_case.CreateUserUseCase
		ListUsersUseCase          list_users.ListUsersUseCase
		DeleteUserUseCase         delete_user.DeleteUserUseCase
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(container_config *ContainerConfig) *Container {
	//grpc injection
	grpc_client := grpc_client.New()
	find_user_service := find_user_service.New(
		pb.NewGetUserServiceClient(
			grpc_client.GetConnection(
				os.Getenv("FIND_USER_SERVICE_URL"))))

	//amqp injection
	amqp_client := amqp_client.New()
	create_user_amqp := create_user_amqp.NewProdocer(amqp_client)

	//integration injection
	http_service := http_service.New()
	json_place_holder_integration := json_place_holder.New(http_service)

	//db injection
	user_repository := model_user.NewUserRepository(container_config.Database)

	//cache injection
	cache_client := cache_client.New()
	users_cache := users_cache.New(cache_client)

	return &Container{
		GetUserGrpcUseCase: get_user_grpc.NewUseCase(find_user_service),
		GetUserUseCase: get_user_use_case.NewUseCase(
			user_repository,
			json_place_holder_integration,
		),
		CreateUserUseCase: create_user_use_case.NewUseCase(
			user_repository,
		),
		DeleteUserUseCase:         delete_user.NewUseCase(user_repository),
		ListUsersUseCase:          list_users.NewUseCase(json_place_holder_integration, users_cache),
		CreateUserProducerUseCase: create_user_producer_use_case.NewUseCase(create_user_amqp),
	}
}
