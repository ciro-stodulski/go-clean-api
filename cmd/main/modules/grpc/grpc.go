package grpc

import (
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/main/modules"
	"go-clean-api/cmd/shared/env"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Engine    *grpc.Server
	container *container.Container
}

func New(c *container.Container) modules.Module {
	return &GRPCServer{container: c, Engine: grpc.NewServer()}
}

func (grpcs *GRPCServer) RunGo() bool {
	return true
}

func (grpcs *GRPCServer) Start() error {
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

	return err
}

func (grpcs *GRPCServer) Stop() {}
