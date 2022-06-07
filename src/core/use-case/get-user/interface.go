package getuserusecase

import (
	entity "go-api/src/core/entities/user"
)

type (
	GetUserUseCase interface {
		GetUser(id string) (*entity.User, error)
	}
)
