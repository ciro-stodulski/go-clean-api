package grpc

import (
	"go-api/src/main/container"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Engine    *grpc.Server
	container *container.Container
}

func (server *GRPCServer) New(container *container.Container) IGrpcServer {
	return &GRPCServer{container: container, Engine: grpc.NewServer()}
}

func (server *GRPCServer) Start() {
	server.LoadServices(server.container)

	reflection.Register(server.Engine)

	// add env para host do grpc
	server_tcp, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print("gRPC: Started with succeffully")

	if err := server.Engine.Serve(server_tcp); err != nil {
		log.Default().Print("Failed to start gRPC server", err)
	}
}
