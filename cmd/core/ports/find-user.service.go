package ports

import (
	entity "go-api/cmd/core/entities/user"
)

type (
	GetUserService interface {
		GetUser(id string) (*entity.User, error)
	}
)
