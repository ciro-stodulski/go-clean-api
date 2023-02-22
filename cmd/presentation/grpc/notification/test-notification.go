package notificaitonpb

import (
	"context"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/grpc/notification/pb"
	"log"
)

type notificationPB struct {
	container *container.Container
}

func New(c *container.Container) *notificationPB {
	return &notificationPB{container: c}
}

func (npb *notificationPB) Verify(ctx context.Context, req *pb.ResquestNotification) (*pb.ResponseNotificaiton, error) {

	log.Default().Println("----> request by grpc: Name:" + req.List.Name + " Describe: " + req.List.Describe)

	npb.container.ListUsersUseCase.ListUsers()

	return &pb.ResponseNotificaiton{
		Event: &pb.List{
			Name:     "succeffully",
			Describe: "grpc connection succeffully",
		},
	}, nil
}