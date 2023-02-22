package domainusecases

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entities/user"
)

type (
	RegisterUserUseCase interface {
		Register(dto domaindto.Dto) (*user.User, error)
	}
)