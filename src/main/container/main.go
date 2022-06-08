package container

import (
	create_user_use_case "go-api/src/core/use-case/create-user"
	create_user_producer_use_case "go-api/src/core/use-case/create-user-producer"
	delete_user "go-api/src/core/use-case/delete-user"
	get_user_use_case "go-api/src/core/use-case/get-user"
	get_user_grpc "go-api/src/core/use-case/get-user-grpc"
	list_users "go-api/src/core/use-case/list-user"
	"os"

	users_cache "go-api/src/infra/cache/users"
	create_user_amqp "go-api/src/infra/integrations/amqp/producer/user-create"
	json_place_holder "go-api/src/infra/integrations/http/jsonplaceholder"
	model_user "go-api/src/infra/repositories/user"

	amqp_client "go-api/src/infra/integrations/amqp/client"
	grpc_client "go-api/src/infra/integrations/grpc/client"
	find_user_service "go-api/src/infra/integrations/grpc/user/get-user"
	"go-api/src/infra/integrations/grpc/user/get-user/pb"
	http_service "go-api/src/infra/integrations/http/client"
	cache_client "go-api/src/main/module/cache/redis"

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
	grpc_client := grpc_client.New()
	find_user_service := find_user_service.New(
		pb.NewGetUserServiceClient(
			grpc_client.GetConnection(
				os.Getenv("FIND_USER_SERVICE_URL"))))

	amqp_client := amqp_client.New()
	create_user_amqp := create_user_amqp.New(amqp_client)

	http_service := http_service.New()
	json_place_holder_integration := json_place_holder.New(http_service)

	user_repository := model_user.NewUserRepository(container_config.Database)

	cache_client := cache_client.New()
	users_cache := users_cache.New(cache_client)

	return &Container{
		GetUserGrpcUseCase: get_user_grpc.New(find_user_service),
		GetUserUseCase: get_user_use_case.New(
			user_repository,
			json_place_holder_integration,
		),
		CreateUserUseCase: create_user_use_case.New(
			user_repository,
		),
		DeleteUserUseCase:         delete_user.New(user_repository),
		ListUsersUseCase:          list_users.New(json_place_holder_integration, users_cache),
		CreateUserProducerUseCase: create_user_producer_use_case.New(create_user_amqp),
	}
}
