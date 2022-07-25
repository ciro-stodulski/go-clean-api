package registeruserusecase

import (
	"go-api/cmd/core/entities/user"
)

func (cuuc *registerUserUseCase) Register(dto Dto) (*user.User, error) {
	u, err := user.New(dto.Email, dto.Password, dto.Name)

	if err != nil {
		return nil, err
	}

	new_u, err := cuuc.UserService.Register(u)

	return new_u, err
}
