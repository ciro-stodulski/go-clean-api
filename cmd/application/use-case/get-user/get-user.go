package getuserusecase

import (
	"go-clean-api/cmd/domain/entity/user"
	portsservice "go-clean-api/cmd/domain/service"
	domainusecases "go-clean-api/cmd/domain/use-case"
)

type (
	getUserUseCase struct {
		UserService portsservice.UserService
	}
)

func New(us portsservice.UserService) domainusecases.GetUserUseCase {
	return &getUserUseCase{
		UserService: us,
	}
}

func (guuc *getUserUseCase) GetUser(id string) (*user.User, error) {
	u, err := guuc.UserService.GetUser(id)

	return u, err
}
