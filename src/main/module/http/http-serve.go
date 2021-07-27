package http_server

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	v1_user "go-api/src/presentation/http/controllers/v1/users"
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

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.NewUserController(container),
	}
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
		api_group := api.Group(ctr.PathGroup())

		for _, route := range ctr.LoadRoutes() {

			switch route.Method {
			case "get":
				api_group.GET(route.Path, route.Function)
			case "post":
				api_group.POST(route.Path, route.Function)
			case "put":
				api_group.PUT(route.Path, route.Function)
			case "patch":
				api_group.PATCH(route.Path, route.Function)
			case "delete":
				api_group.DELETE(route.Path, route.Function)
			default:
			}
		}
	}

	api.GET("/status", func(context *gin.Context) {
		context.Status(http.StatusAccepted)
	})

	gin.SetMode(gin.ReleaseMode)

	api.Use(gin.Recovery())
}
