package container

import (
	"go-api/src/core/useCases/user"
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

func NewContainer(cfg *ContainerConfig) *Container {
	user_ctx := user.NewService(
		model_user.NewUserModel(cfg.Database),
	)

	return &Container{
		UserService: *user.NewService(user_ctx.RepositoryUser),
	}
}
