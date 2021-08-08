package ports

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
)

type (
	Repository interface {
		GetById(id entity_root.ID) (*entity.User, error)
	}
)
