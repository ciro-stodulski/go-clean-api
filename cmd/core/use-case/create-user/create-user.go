package createuserusecase

import (
	"go-api/cmd/core/entities/user"
	dto "go-api/cmd/interface/amqp/consumers/users/create/dto"

	"github.com/google/uuid"
)

func (cuuc *createUserUseCase) CreateUser(dto dto.CreateDto) (*user.User, error) {

	u, err := cuuc.RepositoryUser.GetByEmail(dto.Email)

	if err != nil {
		return nil, err
	}

	if u.ID != uuid.Nil {
		return nil, user.ErrUserAlreadyExists
	}

	new_u, err := user.New(dto.Email, dto.Password, dto.Name)

	if err != nil {
		return nil, err
	}

	cuuc.RepositoryUser.Create(new_u)

	return new_u, err
}
