package getusergrpcusecase

import (
	"go-api/cmd/core/ports"
)

type getUserGrpcUseCase struct {
	GetUserService ports.GetUserService
}

func New(
	gus ports.GetUserService,
) GetUserGrpcUseCase {
	return &getUserGrpcUseCase{
		GetUserService: gus,
	}
}
