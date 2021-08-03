package container

import (
	user "go-api/src/core/useCases/user/get-user"
	jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder"
	model_user "go-api/src/infra/repositories/user"
	http_service "go-api/src/main/module/http/client"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		UserService user.UseCase
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(container_config *ContainerConfig) *Container {

	return &Container{
		UserService: user.NewService(
			model_user.NewUserRepository(container_config.Database),
			jsonplaceholder.New(http_service.New()),
		),
	}
}
