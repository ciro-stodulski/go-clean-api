package createuserproducerusecase

import (
	port "go-api/src/core/ports"
)

type (
	CreateUserUseCase interface {
		CreateUser(dto port.CreateDto) error
	}
)
