package getusergrpcusecase

import (
	getuserservice "go-api/cmd/infra/integrations/grpc/user/get-user"
)

type getUserGrpcUseCase struct {
	GetUserService getuserservice.GetUserService
}

func New(
	gus getuserservice.GetUserService,
) GetUserGrpcUseCase {
	return &getUserGrpcUseCase{
		GetUserService: gus,
	}
}
