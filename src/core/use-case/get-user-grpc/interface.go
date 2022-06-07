package getusergrpcusecase

import (
	"go-api/src/core/entities/user"
)

type (
	GetUserGrpcUseCase interface {
		GetUser(id string) (*user.User, error)
	}
)
