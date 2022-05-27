package grpc_client

import (
	"log"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	Engine *grpc.ClientConn
}

func New(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error during connection ")
	}

	defer connection.Close()

	return connection
}
