package getuserusecase

import (
	"go-clean-api/cmd/domain/entities/user"
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-cases"
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
