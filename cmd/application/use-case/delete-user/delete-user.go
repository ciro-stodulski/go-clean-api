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

func New(us service.UserService) usecase.UseCase[string, interface{}] {
	return &deleteUserUseCase{
		UserService: us,
	}
}

func (duuc *deleteUserUseCase) Perform(id string) (interface{}, error) {
	err := duuc.UserService.DeleteUser(id)

	return nil, err
}
