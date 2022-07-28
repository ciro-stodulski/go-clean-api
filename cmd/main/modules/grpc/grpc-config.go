package grpc

import (
	notificationpb "go-api/cmd/interface/grpc/user"
	"go-api/cmd/interface/grpc/user/pb"
	"go-api/cmd/main/container"
	"log"
)

func (s *GRPCServer) LoadServices(c *container.Container) {
	pb.RegisterNotificationInterfacePbServer(s.Engine, notificationpb.New(c))
	log.Default().Print("gRPC: Services registered")
}
