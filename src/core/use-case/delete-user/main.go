package deleteuserusecase

import (
	"go-api/src/core/ports"
)

type deleteUserUseCase struct {
	RepositoryUser ports.UserRepository
}

func New(ur ports.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		RepositoryUser: ur,
	}
}
