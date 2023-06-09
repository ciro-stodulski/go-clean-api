package grpc

import (
	"go-clean-api/cmd/main/container"
	notificaitonpb "go-clean-api/cmd/presentation/grpc/notification"
	"go-clean-api/cmd/presentation/grpc/notification/pb"
	"log"
)

func (s *GRPCServer) LoadServices(c *container.Container) {
	pb.RegisterTestNotificationInterfacePbServer(s.Engine, notificaitonpb.New(c.ListUsersUseCase))
	log.Default().Print("gRPC: Services registered")
}
