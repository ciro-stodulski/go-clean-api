package find_service

import (
	"go-api/src/main/container"
)

type FindUserService struct {
	container *container.Container
}

func New(c *container.Container) *FindUserService {
	return &FindUserService{container: c}
}
