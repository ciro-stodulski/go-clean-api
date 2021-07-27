package v1_user

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"

	"github.com/gin-gonic/gin"
)

type (
	createController struct {
		container *container.Container
	}

	CreateController interface {
		controllers.Controller
		findById(ctx *gin.Context)
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
			Method:   "get",
			Path:     "/:id",
			Function: createController.findById,
		},
	}
}
