package get_user_grpc

import (
	entity "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

func (service *getUserGrpcUseCase) GetUser(id string) (*entity.User, error) {
	user, err := service.GetUserService.GetUser(id)

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		log.Default().Print("not found user with id:" + id)
		return nil, entity.ErrUserNotFound
	}

	return user, err
}
