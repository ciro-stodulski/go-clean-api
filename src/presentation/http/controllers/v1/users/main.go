package v1_user

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	"go-api/src/presentation/http/middlewares"
	ports_http "go-api/src/presentation/http/ports"
)

type (
	createController struct {
		container *container.Container
	}

	CreateController interface {
		controllers.Controller
		findById(req ports_http.HttpRequest)
	}
)

func NewUserController(c *container.Container) controllers.Controller {
	return &createController{c}
}

func (createController *createController) PathGroup() string {
	return "/v1/users"
}

func (createController *createController) LoadRoutes() []controllers.CreateRoute {
	return []controllers.CreateRoute{
		{
			Method:      "get",
			Path:        "/:id",
			Handle:      createController.findById,
			Middlewares: []controllers.Middleware{middlewares.Log},
		},
	}
}
