package container

import (
	get_user_use_case "go-api/src/core/useCases/get-user"
	list_users "go-api/src/core/useCases/list-user"
	json_place_holder "go-api/src/infra/http/integrations/jsonplaceholder"
	model_user "go-api/src/infra/repositories/user"
	http_service "go-api/src/main/module/http/client"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		GetUserUseCase   get_user_use_case.GetUserUseCase
		ListUsersUseCase list_users.ListUsersUseCase
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(container_config *ContainerConfig) *Container {
	jsonPlaceHolderIntegration := json_place_holder.New(http_service.New(), "https://jsonplaceholder.typicode.com")
	userRepository := model_user.NewUserRepository(container_config.Database)

	return &Container{
		GetUserUseCase: get_user_use_case.NewUseCase(
			userRepository,
			jsonPlaceHolderIntegration,
		),
		ListUsersUseCase: list_users.NewUseCase(jsonPlaceHolderIntegration),
	}
}
