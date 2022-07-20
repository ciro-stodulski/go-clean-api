package getusergrpcusecase

import (
	"go-api/cmd/core/entities/user"
)

type (
	GetUserGrpcUseCase interface {
		GetUser(id string) (*user.User, error)
	}
)
