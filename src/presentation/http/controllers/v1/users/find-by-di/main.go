package v1_user

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	"go-api/src/presentation/http/middlewares"
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
