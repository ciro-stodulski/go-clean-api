package container

import (
	user "go-api/src/core/useCases/user"
	model_user "go-api/src/infra/repositories/user"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		UserService user.Service
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(container_config *ContainerConfig) *Container {
	user_context := user.NewService(
		model_user.NewUserRepository(container_config.Database),
	)

	return &Container{
		UserService: *user.NewService(user_context.RepositoryUser),
	}
}
