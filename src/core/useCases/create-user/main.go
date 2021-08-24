package get_user

import (
	interfaces "go-api/src/core/ports"
)

type createUserUseCase struct {
	RepositoryUser interfaces.UserRepository
}

func NewUseCase(repository interfaces.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		RepositoryUser: repository,
	}
}
