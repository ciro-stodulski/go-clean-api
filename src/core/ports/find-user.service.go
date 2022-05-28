package ports

import (
	entity "go-api/src/core/entities/user"
)

type (
	GetUserService interface {
		GetUser(ID string) (*entity.User, error)
	}
)
