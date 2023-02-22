package domainusersql

import (
	entity_root "go-clean-api/cmd/domain/entities"
	entity "go-clean-api/cmd/domain/entities/user"
)

type (
	UserSql interface {
		DeleteById(id entity_root.ID) error
		GetById(id entity_root.ID) (*entity.User, error)
		GetByEmail(email string) (*entity.User, error)
		Create(u *entity.User) error
	}
)
