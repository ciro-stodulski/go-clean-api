package user_use_case

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
)

type (
	//Repository interface
	Repository interface {
		GetById(id entity_root.ID) (*entity.User, error)
	}

	Integration interface {
		GetTodos() error
	}

	//UseCase interface
	UseCase interface {
		GetUser(id string) (*entity.User, error)
	}
)
