package get_user_grpc

import (
	interfaces "go-api/src/core/ports"
)

type getUserGrpcUseCase struct {
	GetUserService interfaces.GetUserService
}

func NewUseCase(
	getUserService interfaces.GetUserService,
) GetUserGrpcUseCase {
	return &getUserGrpcUseCase{
		GetUserService: getUserService,
	}
}
