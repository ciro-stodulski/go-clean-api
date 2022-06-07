package create_user_producer

import (
	dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
)

type (
	CreateUserUseCase interface {
		CreateUser(dto dto.CreateDto) error
	}
)
