package domainusecases

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
)

type (
	RegisterUserUseCase interface {
		Register(dto domaindto.Dto) (*user.User, *domainexceptions.ApplicationException, error)
	}
)
