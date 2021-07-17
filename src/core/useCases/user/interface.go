package user

import (
	entity "go-api/src/core/entities"
)

//Reader interface
type Reader interface {
	GetById(id entity.ID) (*entity.User, error)
}

//Repository interface
type Repository interface {
	Reader
}

//UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
}
