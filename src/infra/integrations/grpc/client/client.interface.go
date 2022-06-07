package grpcclient

import "google.golang.org/grpc"

type (
	GRPCClient interface {
		GetConnection(host string) *grpc.ClientConn
	}
)
