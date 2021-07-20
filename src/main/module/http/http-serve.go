package http_server

import (
	controllers "go-api/src/apresentation/http/controllers"
	v1_user "go-api/src/apresentation/http/controllers/v1/users"

	"go-api/src/main/container"

	"github.com/gin-gonic/gin"
)

func _loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.NewUserController(container),
	}
}

func SetupRoutes(gin *gin.Engine, c *container.Container) {
	api := gin.Group("/")

	controls := _loadControllers(c)

	for _, ctr := range controls {
		ctr.Register(api)
	}

	api.GET("/status", controllers.HealthCheck)
}
