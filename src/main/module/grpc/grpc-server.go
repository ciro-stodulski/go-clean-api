package grpc

import (
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

	loadServices(server.Engine)

	reflection.Register(server.Engine)

	server_tcp, err := net.Listen("tcp", ":50055")

	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print("gRPC: Started with succeffully")

	if err := server.Engine.Serve(server_tcp); err != nil {
		log.Default().Print("Failed to start gRPC server", err)
	}

	return nil
}
