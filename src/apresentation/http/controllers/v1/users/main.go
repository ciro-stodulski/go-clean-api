package v1_user

import (
	controllers "go-api/src/apresentation/http/controllers"
	"go-api/src/main/container"

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

func (createController *createController) Register(gr *gin.RouterGroup) {
	user_group := gr.Group("/v1/users")

	user_group.GET("/:id", createController.findById)
}
