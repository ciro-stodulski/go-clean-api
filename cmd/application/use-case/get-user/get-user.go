package getuserusecase

import (
	"go-clean-api/cmd/domain/entity/user"
	portsservice "go-clean-api/cmd/domain/service"
	usecase "go-clean-api/cmd/domain/use-case"
)

type (
	getUserUseCase struct {
		UserService portsservice.UserService
	}
)

func New(us portsservice.UserService) usecase.UseCase[string, *user.User] {
	return &getUserUseCase{
		UserService: us,
	}
}

func (guuc *getUserUseCase) Perform(id string) (*user.User, error) {
	u, err := guuc.UserService.GetUser(id)

	return u, err
}
