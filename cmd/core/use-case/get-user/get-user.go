package getuserusecase

import (
	"go-api/cmd/core/entities/user"
	portsservice "go-api/cmd/core/ports-service"
)

type getUserUseCase struct {
	UserService portsservice.UserService
}

func New(us portsservice.UserService) GetUserUseCase {
	return &getUserUseCase{
		UserService: us,
	}
}

func (guuc *getUserUseCase) GetUser(id string) (*user.User, error) {
	u, err := guuc.UserService.GetUser(id)

	return u, err
}
