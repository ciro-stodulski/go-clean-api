package createuserproducerusecase

import (
	port "go-api/cmd/core/ports"
)

type (
	CreateUserUseCase interface {
		CreateUser(dto port.CreateDto) error
	}
)
