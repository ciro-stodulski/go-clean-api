package getusergrpcusecase

import (
	user "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

func (guguc *getUserGrpcUseCase) GetUser(id string) (*user.User, error) {
	u, err := guguc.GetUserService.GetUser(id)

	if err != nil {
		return nil, err
	}

	if u.ID == uuid.Nil {
		log.Default().Print("not found user with id:" + id)
		return nil, user.ErrUserNotFound
	}

	return u, err
}
