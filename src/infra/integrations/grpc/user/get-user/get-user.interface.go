package getuserservice

import (
	"context"
	"go-api/src/infra/integrations/grpc/user/get-user/pb"

	"google.golang.org/grpc"
)

type GetUserService interface {
	GetUser(context.Context, *pb.NewRequestGetUser, ...grpc.CallOption) (*pb.NewResponseGetUser, error)
}
