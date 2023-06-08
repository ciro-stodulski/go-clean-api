package deleteuserusecase

import (
	"go-clean-api/cmd/domain/service"
	usecase "go-clean-api/cmd/domain/use-case"
)

type (
	deleteUserUseCase struct {
		UserService service.UserService
	}
)

func New(us service.UserService) usecase.DeleteUserUseCase {
	return &deleteUserUseCase{
		UserService: us,
	}
}

func (duuc *deleteUserUseCase) DeleteUser(id string) error {
	err := duuc.UserService.DeleteUser(id)

	return err
}
