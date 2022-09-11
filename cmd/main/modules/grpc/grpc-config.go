package grpc

import (
	notificaitonpb "go-clean-api/cmd/interface/grpc/notification"
	"go-clean-api/cmd/interface/grpc/notification/pb"
	"go-clean-api/cmd/main/container"
	"log"
)

func (s *GRPCServer) LoadServices(c *container.Container) {
	pb.RegisterTestNotificationInterfacePbServer(s.Engine, notificaitonpb.New(c))
	log.Default().Print("gRPC: Services registered")
}
