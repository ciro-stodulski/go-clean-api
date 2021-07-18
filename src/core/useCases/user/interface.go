package user

import (
	entity "go-api/src/core/entities"
)

//Repository interface
type Repository interface {
	GetById(id entity.ID) (*entity.User, error)
}

//UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
}
