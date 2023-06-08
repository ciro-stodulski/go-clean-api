package usecase

import (
	"go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"
)

type (
	RegisterUserUseCase interface {
		Register(dto dto.RegisterUser) (*user.User, error)
	}
)
