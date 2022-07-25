package getuserservice

import (
	"context"
	entity "go-api/cmd/core/entities/user"
	"go-api/cmd/infra/integrations/grpc/user/get-user/pb"

	"google.golang.org/grpc"
)

type (
	PbGetUserService interface {
		GetUser(context.Context, *pb.NewRequestGetUser, ...grpc.CallOption) (*pb.NewResponseGetUser, error)
	}

	GetUserService interface {
		GetUser(id string) (*entity.User, error)
	}
)
