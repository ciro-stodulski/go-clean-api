package get_user

import (
	entity "go-api/src/core/entities/user"
)

type (
	UseCase interface {
		GetUser(id string) (*entity.User, error)
	}
)
