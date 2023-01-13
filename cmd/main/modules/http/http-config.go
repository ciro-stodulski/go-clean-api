package http

import (
	controllers "go-clean-api/cmd/interface/http/controllers"
	v1_user "go-clean-api/cmd/interface/http/controllers/v1/users/find-by-di"
	"go-clean-api/cmd/interface/http/middlewares"
	"go-clean-api/cmd/main/container"
)

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.New(container),
	}
}

func loadMiddlewaresGlobals() []controllers.Middleware {
	return []controllers.Middleware{
		middlewares.Global,
	}
}
