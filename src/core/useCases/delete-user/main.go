package delete_user

import (
	interfaces "go-api/src/core/ports"
)

type deleteUserUseCase struct {
	RepositoryUser interfaces.UserRepository
}

func NewUseCase(repository interfaces.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		RepositoryUser: repository,
	}
}
