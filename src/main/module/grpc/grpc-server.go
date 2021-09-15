package grpc

import (
	finduserservice "go-api/src/presentation/grpc/services/user/find-user"
	"go-api/src/presentation/grpc/services/user/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Engine *grpc.Server
}

func (server *GRPCServer) Start() error {
	server.Engine = grpc.NewServer()
	pb.RegisterUserServiceServer(server.Engine, &finduserservice.FindUserService{})
	reflection.Register(server.Engine)

	server_tcp, err := net.Listen("tcp", ":50055")

	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print("gRPC: Started with succeffully")
	server.Engine.Serve(server_tcp)

	return nil
}
