package domainusecases

import (
	"go-clean-api/cmd/domain/entities/user"
)

type (
	GetUserUseCase interface {
		GetUser(id string) (*user.User, error)
	}
)
