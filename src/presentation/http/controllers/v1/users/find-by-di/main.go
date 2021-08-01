package v1_user

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	ports_http "go-api/src/presentation/http/ports"

	"go-api/src/presentation/http/middlewares"
)

type (
	CreateController struct {
		container *container.Container
	}

	ICreateController interface {
		controllers.Controller
		findById(req ports_http.HttpRequest)
	}
)

func NewController(c *container.Container) controllers.Controller {
	return &CreateController{c}
}

func (createController *CreateController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "get",
		Path:        "/:id",
		Handle:      createController.findById,
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}
