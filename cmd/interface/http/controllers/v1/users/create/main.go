package v1_user_create

import (
	"go-api/cmd/core/ports"
	controllers "go-api/cmd/interface/http/controllers"
	"go-api/cmd/main/container"
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
