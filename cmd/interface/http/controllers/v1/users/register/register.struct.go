package controllerv1userregister

import (
	controllers "go-api/cmd/interface/http/controllers"
	"go-api/cmd/main/container"
)

type (
	registerController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &registerController{c}
}
