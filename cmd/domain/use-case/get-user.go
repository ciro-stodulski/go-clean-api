package domainusecases

import (
	"go-clean-api/cmd/domain/entity/user"
)

type (
	GetUserUseCase interface {
		GetUser(id string) (*user.User, error)
	}
)
