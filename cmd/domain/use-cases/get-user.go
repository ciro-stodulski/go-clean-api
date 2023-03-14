package domainusecases

import (
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
)

type (
	GetUserUseCase interface {
		GetUser(id string) (*user.User, *domainexceptions.ApplicationException, error)
	}
)
