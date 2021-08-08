package http_server

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	v1_user "go-api/src/presentation/http/controllers/v1/users/find-by-di"
	"go-api/src/presentation/http/middlewares"
)

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.NewController(container),
	}
}

func loadMiddlewaresGlobals() []controllers.Middleware {
	return []controllers.Middleware{
		middlewares.Global,
	}
}
