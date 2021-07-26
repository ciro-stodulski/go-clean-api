package http_server

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	v1_user "go-api/src/presentation/http/controllers/v1/users"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Engine *gin.Engine
}

type IHttpServer interface {
	Start() error
	New(container *container.Container)
}

func (server *HttpServer) Start() error {
	port := os.Getenv("HOST_PORT")

	log.Default().Print("server started with succeffully")

	return server.Engine.Run("localhost:" + port)
}

func (server *HttpServer) New(container *container.Container) {
	server.Engine = gin.New()

	gin.SetMode(gin.DebugMode)

	server.Engine.Use(gin.Recovery())

	api := server.Engine.Group("/")
	controls := loadControllers(container)

	for _, ctr := range controls {
		ctr.Register(api)
	}

	api.GET("/status", controllers.HealthCheck)
	gin.SetMode(gin.ReleaseMode)

	api.Use(gin.Recovery())
}

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.NewUserController(container),
	}
}
