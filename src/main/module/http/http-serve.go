package http_server

import (
	"go-api/src/main/container"
	"log"
	"net/http"
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
	gin.SetMode(gin.DebugMode)
	server.Engine = gin.New()

	api := server.Engine.Group("/")

	controls := loadControllers(container)
	loadRoutes(controls, *api)

	api.GET("/status", func(context *gin.Context) {
		context.Status(http.StatusAccepted)
	})

	gin.SetMode(gin.ReleaseMode)

	api.Use(gin.Recovery())
}
