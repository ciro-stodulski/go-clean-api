package grpcclient

import (
	"log"

	"google.golang.org/grpc"
)

type (
	GRPCClient interface {
		GetConnection(host string) *grpc.ClientConn
	}

	gRPCClientService struct{}
)

func New() GRPCClient {
	return &gRPCClientService{}
}

func (gclient *gRPCClientService) GetConnection(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error during connection ")
	}

	return connection
}
