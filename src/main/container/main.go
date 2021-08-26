package container

import (
	create_user_use_case "go-api/src/core/useCases/create-user"
	get_user_use_case "go-api/src/core/useCases/get-user"
	list_users "go-api/src/core/useCases/list-user"
	users_cache "go-api/src/infra/cache/users"
	json_place_holder "go-api/src/infra/http/integrations/jsonplaceholder"
	model_user "go-api/src/infra/repositories/user"
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
		GetUserUseCase    get_user_use_case.GetUserUseCase
		CreateUserUseCase create_user_use_case.CreateUserUseCase
		ListUsersUseCase  list_users.ListUsersUseCase
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(container_config *ContainerConfig) *Container {
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
		ListUsersUseCase: list_users.NewUseCase(json_place_holder_integration, users_cache),
	}
}
