package v1_user_grpc

import (
	controllers "go-api/cmd/interface/http/controllers"
	"go-api/cmd/interface/http/middlewares"
	"go-api/cmd/main/container"
)

type (
	findByIdGrpcController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &findByIdGrpcController{c}
}

func (findByIdGrpcController *findByIdGrpcController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/grpc/users",
		Method:      "get",
		Path:        "/:id",
		Handle:      findByIdGrpcController.findById,
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}
