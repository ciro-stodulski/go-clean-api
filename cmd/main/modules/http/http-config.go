package http

import (
	controllers "go-api/cmd/interface/http/controllers"
	v1_user_create "go-api/cmd/interface/http/controllers/v1/users/create"
	v1_user_delete "go-api/cmd/interface/http/controllers/v1/users/delete"
	v1_user "go-api/cmd/interface/http/controllers/v1/users/find-by-di"
	v1_user_grpc "go-api/cmd/interface/http/controllers/v1/users/find-by-id-grpc"
	"go-api/cmd/interface/http/middlewares"
	"go-api/cmd/main/container"
)

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.New(container),
		v1_user_grpc.New((container)),
		v1_user_create.New(container),
		v1_user_delete.New(container),
	}
}

func loadMiddlewaresGlobals() []controllers.Middleware {
	return []controllers.Middleware{
		middlewares.Global,
	}
}
