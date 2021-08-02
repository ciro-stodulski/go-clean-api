package v1_user

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	"go-api/src/presentation/http/middlewares"
)

type (
	createController struct {
		container *container.Container
	}
)

func NewController(c *container.Container) controllers.Controller {
	return &createController{c}
}

func (createController *createController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "get",
		Path:        "/:id",
		Handle:      createController.findById,
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}
