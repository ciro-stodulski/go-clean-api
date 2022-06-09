package v1_user_create

import (
	"go-api/src/core/ports"
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
)

type (
	createController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &createController{c}
}

func (createController *createController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Handle:   createController.create,
		Dto:      ports.CreateDto{},
	}
}
