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

func (server *GRPCServer) New(c *container.Container) IGrpcServer {
	return &GRPCServer{container: c, Engine: grpc.NewServer()}
}

func (grpcs *GRPCServer) Start() {
	grpcs.LoadServices(grpcs.container)

	reflection.Register(grpcs.Engine)

	server, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print("gRPC: Started with succeffully")

	if err := grpcs.Engine.Serve(server); err != nil {
		log.Default().Print("Failed to start gRPC server", err)
	}
}
