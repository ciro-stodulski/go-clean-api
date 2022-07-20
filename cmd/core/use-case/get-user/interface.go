package getuserusecase

import (
	entity "go-api/cmd/core/entities/user"
)

type (
	GetUserUseCase interface {
		GetUser(id string) (*entity.User, error)
	}
)
