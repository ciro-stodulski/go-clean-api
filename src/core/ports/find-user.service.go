package ports

import (
	entity "go-api/src/core/entities/user"
)

type (
	GetUserService interface {
		GetUser(id string) (*entity.User, error)
	}
)
