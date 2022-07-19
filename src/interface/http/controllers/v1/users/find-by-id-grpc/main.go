package v1_user_grpc

import (
	controllers "go-api/src/interface/http/controllers"
	"go-api/src/interface/http/middlewares"
	"go-api/src/main/container"
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
