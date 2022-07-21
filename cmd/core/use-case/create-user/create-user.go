package createuserusecase

import (
	"go-api/cmd/core/entities/user"
	portsservice "go-api/cmd/core/ports-service"
	dto "go-api/cmd/interface/amqp/consumers/users/create/dto"
)

type createUserUseCase struct {
	UserService portsservice.UserService
}

func New(us portsservice.UserService) CreateUserUseCase {
	return &createUserUseCase{
		UserService: us,
	}
}

func (cuuc *createUserUseCase) CreateUser(dto dto.CreateDto) (*user.User, error) {
	new_u, err := cuuc.UserService.CreateUser(dto)

	return new_u, err
}
