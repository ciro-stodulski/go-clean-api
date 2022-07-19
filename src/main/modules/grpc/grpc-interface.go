package grpc

import (
	"go-api/src/main/container"
)

type IGrpcServer interface {
	New(c *container.Container) IGrpcServer
	Start()
	LoadServices(c *container.Container)
}
