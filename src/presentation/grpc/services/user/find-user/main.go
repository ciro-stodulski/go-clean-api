package find_service

import (
	"go-api/src/main/container"
)

type FindUserService struct {
	container *container.Container
}

func NewService(container *container.Container) *FindUserService {
	return &FindUserService{container: container}
}
