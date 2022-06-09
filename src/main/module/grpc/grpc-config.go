package grpc

import (
	"go-api/src/main/container"
	finduserservice "go-api/src/presentation/grpc/services/user/find-user"
	"go-api/src/presentation/grpc/services/user/pb"
	"log"
)

func (s *GRPCServer) LoadServices(c *container.Container) {
	pb.RegisterFindUserServiceServer(s.Engine, finduserservice.New(c))
	log.Default().Print("gRPC: Services registered")
}
