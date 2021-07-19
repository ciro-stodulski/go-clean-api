package container

import (
	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
	}

	Container struct {
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{}
}

func NewContainer(cfg *ContainerConfig) *Container {

	return &Container{}
}
