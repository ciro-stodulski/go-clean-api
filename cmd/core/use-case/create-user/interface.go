package createuserusecase

import (
	entity "go-api/cmd/core/entities/user"
	dto "go-api/cmd/interface/amqp/consumers/users/create/dto"
)

type (
	CreateUserUseCase interface {
		CreateUser(dto dto.CreateDto) (*entity.User, error)
	}
)
