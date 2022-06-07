package getuserservice

import (
	"context"
	ports "go-api/src/core/ports"
	"go-api/src/infra/integrations/grpc/user/get-user/pb"

	"google.golang.org/grpc"
)

type GetUserService interface {
	GetUser(context.Context, *pb.NewRequestGetUser, ...grpc.CallOption) (*pb.NewResponseGetUser, error)
}

type getUserService struct {
	service GetUserService
}

func New(service GetUserService) ports.GetUserService {

	return &getUserService{
		service: service,
	}
}
