package createuserusecase

import (
	"go-api/src/core/entities/user"
	dto "go-api/src/presentation/amqp/consumers/users/create/dto"

	"github.com/google/uuid"
)

func (usecase *createUserUseCase) CreateUser(dto dto.CreateDto) (*user.User, error) {

	u, err := usecase.RepositoryUser.GetByEmail(dto.Email)

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

	usecase.RepositoryUser.Create(new_u)

	return new_u, err
}
