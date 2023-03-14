package deleteuserusecase

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
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

func (duuc *deleteUserUseCase) DeleteUser(id string) (*domainexceptions.ApplicationException, error) {
	errApp, err := duuc.UserService.DeleteUser(id)

	return errApp, err
}
