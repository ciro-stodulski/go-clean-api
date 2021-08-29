package container

import (
	create_user_use_case "go-api/src/core/useCases/create-user"
	create_user_producer_use_case "go-api/src/core/useCases/create-user-producer"
	get_user_use_case "go-api/src/core/useCases/get-user"
	list_users "go-api/src/core/useCases/list-user"

	create_user_amqp "go-api/src/infra/amqp/producer/user-create"
	users_cache "go-api/src/infra/cache/users"
	json_place_holder "go-api/src/infra/http/integrations/jsonplaceholder"
	model_user "go-api/src/infra/repositories/user"

	amqp_client "go-api/src/main/module/amqp/rabbitmq/client"
	cache_client "go-api/src/main/module/cache/redis"
	http_service "go-api/src/main/module/http/client"
	"os"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		GetUserUseCase            get_user_use_case.GetUserUseCase
		CreateUserUseCase         create_user_use_case.CreateUserUseCase
		CreateUserProducerUseCase create_user_producer_use_case.CreateUserUseCase
		ListUsersUseCase          list_users.ListUsersUseCase
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(container_config *ContainerConfig) *Container {
	//amqp injection
	amqp_client := amqp_client.New()
	create_user_amqp := create_user_amqp.NewProdocer(amqp_client)

	//integration injection
	http_service := http_service.New()
	json_place_holder_integration := json_place_holder.New(http_service, os.Getenv("JSON_PLACE_OLDER_INTEGRATION_URL"))

	//db injection
	user_repository := model_user.NewUserRepository(container_config.Database)

	//cache injection
	cache_client := cache_client.New()
	users_cache := users_cache.New(cache_client)

	return &Container{
		GetUserUseCase: get_user_use_case.NewUseCase(
			user_repository,
			json_place_holder_integration,
		),
		CreateUserUseCase: create_user_use_case.NewUseCase(
			user_repository,
		),
		ListUsersUseCase:          list_users.NewUseCase(json_place_holder_integration, users_cache),
		CreateUserProducerUseCase: create_user_producer_use_case.NewUseCase(create_user_amqp),
	}
}
