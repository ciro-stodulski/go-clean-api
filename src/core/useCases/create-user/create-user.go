package get_user

import (
	entity "go-api/src/core/entities/user"
	dto "go-api/src/presentation/amqp/consumers/users/create/dto"

	"github.com/google/uuid"
)

func (service *createUserUseCase) CreateUser(dto dto.CreateDto) (*entity.User, error) {

	user, err := service.RepositoryUser.GetByEmail(dto.Email)

	if err != nil {
		return nil, err
	}

	if user.ID != uuid.Nil {
		return nil, entity.ErrUserAlreadyExists
	}

	new_user, err := entity.NewUser(dto.Email, dto.Password, dto.Name)

	if err != nil {
		return nil, err
	}

	service.RepositoryUser.Create(*new_user)

	return new_user, err
}
