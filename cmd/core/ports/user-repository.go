package ports

import (
	entity_root "go-api/cmd/core/entities"
	entity "go-api/cmd/core/entities/user"
)

type (
	UserRepository interface {
		DeleteById(id entity_root.ID) error
		GetById(id entity_root.ID) (*entity.User, error)
		GetByEmail(email string) (*entity.User, error)
		Create(u *entity.User)
	}
)
