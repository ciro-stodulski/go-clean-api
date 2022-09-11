package deleteuserusecase

import (
	portsservice "go-clean-api/cmd/core/ports"
)

type (
	DeleteUserUseCase interface {
		DeleteUser(id string) error
	}
	deleteUserUseCase struct {
		UserService portsservice.UserService
	}
)

func New(us portsservice.UserService) DeleteUserUseCase {
	return &deleteUserUseCase{
		UserService: us,
	}
}

func (duuc *deleteUserUseCase) DeleteUser(id string) (err error) {
	duuc.UserService.DeleteUser(id)

	return
}
