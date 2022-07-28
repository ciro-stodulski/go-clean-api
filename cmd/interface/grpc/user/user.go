package userpb

import (
	"context"
	"go-api/cmd/interface/grpc/user/pb"
	"go-api/cmd/main/container"
	"log"
)

type userPB struct {
	container *container.Container
}

func New(c *container.Container) *userPB {
	return &userPB{container: c}
}

func (npb *userPB) Verify(ctx context.Context, req *pb.ResquestUser) (*pb.ResponseUser, error) {

	log.Default().Println("----> request by grpc: Name:" + req.List.Name + " Describe: " + req.List.Describe)

	npb.container.ListUsersUseCase.ListUsers()

	return &pb.ResponseUser{
		Event: &pb.List{
			Name:     "succeffully",
			Describe: "grpc connection succeffully",
		},
	}, nil
}
