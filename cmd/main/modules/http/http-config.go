package http

import (
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/http/controller"
	deleteuser "go-clean-api/cmd/presentation/http/controller/v1/users/delete"
	findbyiduser "go-clean-api/cmd/presentation/http/controller/v1/users/find-by-di"
	controllerv1userregister "go-clean-api/cmd/presentation/http/controller/v1/users/register"
	"go-clean-api/cmd/presentation/http/middlewares"
)

func loadControllers(container *container.Container) []controller.Controller {
	return []controller.Controller{
		controllerv1userregister.New(container.RegisterUserUseCase),
		findbyiduser.New(container.GetUserUseCase),
		deleteuser.New(container.DeleteUserUseCase),
	}
}

func loadMiddlewaresGlobals() []controller.Middleware {
	return []controller.Middleware{
		middlewares.Global,
	}
}
