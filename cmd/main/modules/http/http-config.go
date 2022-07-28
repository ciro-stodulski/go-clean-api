package http

import (
	controllers "go-api/cmd/interface/http/controllers"
	v1_user_delete "go-api/cmd/interface/http/controllers/v1/users/delete"
	v1_user "go-api/cmd/interface/http/controllers/v1/users/find-by-di"
	controllerv1userregister "go-api/cmd/interface/http/controllers/v1/users/register"
	"go-api/cmd/interface/http/middlewares"
	"go-api/cmd/main/container"
)

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		controllerv1userregister.New(container),
		v1_user.New(container),
		v1_user_delete.New(container),
	}
}

func loadMiddlewaresGlobals() []controllers.Middleware {
	return []controllers.Middleware{
		middlewares.Global,
	}
}
