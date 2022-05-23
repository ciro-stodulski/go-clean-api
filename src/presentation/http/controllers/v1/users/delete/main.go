package v1_delete_user

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
		Method:      "delete",
		Path:        "/:id",
		Handle:      createController.deleteById,
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}
