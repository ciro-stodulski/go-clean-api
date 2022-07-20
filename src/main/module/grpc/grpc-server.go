package grpc

import (
	"go-api/src/main/container"
	"go-api/src/shared/env"
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

	server, err := net.Listen("tcp", ":"+env.Env().GrpcServerPort)

	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print("gRPC: Started with succeffully in port: " + env.Env().GrpcServerPort)

	if err := grpcs.Engine.Serve(server); err != nil {
		log.Default().Print("Failed to start gRPC server", err)
	}
}
