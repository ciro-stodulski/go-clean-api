package grpc_client

import (
	"log"

	"google.golang.org/grpc"
)

type (
	gRPCClientService struct{}
)

func New() GRPCClient {
	return &gRPCClientService{}
}

func (gClient *gRPCClientService) GetConnection(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error during connection ")
	}

	return connection
}
