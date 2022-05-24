package grpc

import (
	finduserservice "go-api/src/presentation/grpc/services/user/find-user"
	"go-api/src/presentation/grpc/services/user/find-user/pb"
	"log"

	"google.golang.org/grpc"
)

func loadServices(engine *grpc.Server) {
	pb.RegisterFindUserServiceServer(engine, &finduserservice.FindUserService{})
	log.Default().Print("gRPC: Services registered")
}
