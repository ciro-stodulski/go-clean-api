package v1_user

import (
	controllers "go-api/cmd/interface/http/controllers"
	"go-api/cmd/interface/http/middlewares"
	"go-api/cmd/main/container"
)

type (
	findByIdController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &findByIdController{c}
}

func (findByIdController *findByIdController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "get",
		Path:        "/:id",
		Handle:      findByIdController.findById,
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}
