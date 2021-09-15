package finduserservice

import (
	"context"
	"go-api/src/presentation/grpc/services/user/pb"
)

type FindUserService struct {
}

func (find_user_service *FindUserService) FindUser(ctx context.Context, req *pb.NewRequestFindUser) (*pb.NewResponseFindUser, error) {

	return &pb.NewResponseFindUser{
		User: &pb.User{
			Name:  "test",
			Email: "test",
		},
	}, nil
}
