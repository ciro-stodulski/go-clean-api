package notificaitonpb

import (
	"context"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/grpc/notification/pb"
	"log"
)

type notificationPB struct {
	ListUsersUseCase usecase.IUseCase[interface{}, interface{}]
}

func New(ListUsersUseCase usecase.IUseCase[interface{}, interface{}]) *notificationPB {
	return &notificationPB{ListUsersUseCase}
}

func (npb *notificationPB) Verify(ctx context.Context, req *pb.ResquestNotification) (*pb.ResponseNotificaiton, error) {

	log.Default().Println("----> request by grpc: Name:" + req.List.Name + " Describe: " + req.List.Describe)

	npb.ListUsersUseCase.Perform(nil)

	return &pb.ResponseNotificaiton{
		Event: &pb.List{
			Name:     "succeffully",
			Describe: "grpc connection succeffully",
		},
	}, nil
}
