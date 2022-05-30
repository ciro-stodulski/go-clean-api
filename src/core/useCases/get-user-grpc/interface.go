package get_user_grpc

import (
	entity "go-api/src/core/entities/user"
)

type (
	GetUserGrpcUseCase interface {
		GetUser(id string) (*entity.User, error)
	}
)
