package createuserusecase

import (
	entity "go-api/src/core/entities/user"
	dto "go-api/src/presentation/amqp/consumers/users/create/dto"
)

type (
	CreateUserUseCase interface {
		CreateUser(dto dto.CreateDto) (*entity.User, error)
	}
)
