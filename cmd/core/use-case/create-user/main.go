package createuserusecase

import (
	"go-api/cmd/core/ports"
)

type createUserUseCase struct {
	RepositoryUser ports.UserRepository
}

func New(ur ports.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		RepositoryUser: ur,
	}
}
