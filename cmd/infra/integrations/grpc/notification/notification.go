package notificationpbgrpc

import (
	"context"
	"fmt"
	"go-api/cmd/infra/integrations/grpc/notification/pb"
	"log"

	"google.golang.org/grpc"
)

type (
	PbNotification interface {
		Verify(context.Context, *pb.Request, ...grpc.CallOption) (*pb.Reponse, error)
	}

	NotificationPbGrpc interface {
		Verify(id string) error
	}

	notificationPbGrpc struct {
		service PbNotification
	}
)

func New(service PbNotification) NotificationPbGrpc {

	return &notificationPbGrpc{
		service: service,
	}
}

func (npbgrpc *notificationPbGrpc) Verify(msg string) error {
	req := &pb.Request{
		Msg: msg,
	}

	res, err := npbgrpc.service.Verify(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Default().Println(res)

	return nil
}
