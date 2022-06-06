package grpc

import (
	"go-api/src/main/container"
	finduserservice "go-api/src/presentation/grpc/services/user/find-user"
	"go-api/src/presentation/grpc/services/user/pb"
	"log"
)

func (server *GRPCServer) LoadServices(container *container.Container) {
	pb.RegisterFindUserServiceServer(server.Engine, finduserservice.NewService(container))
	log.Default().Print("gRPC: Services registered")
}
