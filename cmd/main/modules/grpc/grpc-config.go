package grpc

import (
	finduserservice "go-api/cmd/interface/grpc/services/user/find-user"
	"go-api/cmd/interface/grpc/services/user/pb"
	"go-api/cmd/main/container"
	"log"
)

func (s *GRPCServer) LoadServices(c *container.Container) {
	pb.RegisterFindUserServiceServer(s.Engine, finduserservice.New(c))
	log.Default().Print("gRPC: Services registered")
}
