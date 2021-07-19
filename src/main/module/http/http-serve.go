package http_server

import (
	controllers "go-api/src/apresentation/http"
	"go-api/src/main/container"

	"github.com/gin-gonic/gin"
)

func _loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{}
}

func SetupRoutes(e *gin.Engine, c *container.Container) {
	api := e.Group("/")

	controls := _loadControllers(c)

	for _, ctr := range controls {
		ctr.Register(api)
	}

	api.GET("/status", controllers.HealthCheck)
}
