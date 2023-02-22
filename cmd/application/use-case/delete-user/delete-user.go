package deleteuserusecase

import (
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-cases"
)

type (
	deleteUserUseCase struct {
		UserService portsservice.UserService
	}
)

func New(us portsservice.UserService) domainusecases.DeleteUserUseCase {
	return &deleteUserUseCase{
		UserService: us,
	}
}

func (duuc *deleteUserUseCase) DeleteUser(id string) (err error) {
	duuc.UserService.DeleteUser(id)

	return
}
