package v1_delete_user

import (
	controllers "go-api/cmd/interface/http/controllers"
	"go-api/cmd/interface/http/middlewares"
	"go-api/cmd/main/container"
)

type (
	deleteController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &deleteController{c}
}

func (deleteController *deleteController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "delete",
		Path:        "/:id",
		Handle:      deleteController.deleteById,
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}
