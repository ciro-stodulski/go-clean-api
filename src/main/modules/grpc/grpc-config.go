package grpc

import (
	finduserservice "go-api/src/interface/grpc/services/user/find-user"
	"go-api/src/interface/grpc/services/user/pb"
	"go-api/src/main/container"
	"log"
)

func (s *GRPCServer) LoadServices(c *container.Container) {
	pb.RegisterFindUserServiceServer(s.Engine, finduserservice.New(c))
	log.Default().Print("gRPC: Services registered")
}
