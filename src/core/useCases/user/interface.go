package user_use_case

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
)

//Repository interface
type Repository interface {
	GetById(id entity_root.ID) (*entity.User, error)
}

//UseCase interface
type UseCase interface {
	GetUser(id entity_root.ID) (*entity.User, error)
}
