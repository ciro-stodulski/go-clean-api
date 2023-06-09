package http

import (
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/main/modules"
	"go-clean-api/cmd/presentation/http/controller"
	"go-clean-api/cmd/shared/env"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type httpServer struct {
	Engine *gin.Engine
}

func (server *httpServer) Start() error {
	log.Default().Print("Http: Server started with succeffully")

	return server.Engine.Run(env.Env().HostHttp + ":" + env.Env().HostPort)
}

func (server *httpServer) Stop() {
}

func (server *httpServer) RunGo() bool {
	return false
}

//	@title			Go architecture APIs
//	@version		1.0
//	@description	Testing Swagger APIs.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						token

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/api/v1

//	@schemes	http
func New(container *container.Container) modules.Module {
	server := &httpServer{}

	gin.SetMode(gin.DebugMode)
	server.Engine = gin.New()

	api := server.Engine.Group("/")

	controls := loadControllers(container)

	if len(loadMiddlewaresGlobals()) > 0 {
		for _, middleware := range loadMiddlewaresGlobals() {
			mds := func(context *gin.Context) {
				params := loadParams(context)
				middleware(controller.HttpRequest{
					Params:  params,
					Query:   context.Request.URL.Query(),
					Headers: context.Request.Header,
					Next:    context.Next,
				})
			}
			api.Use(mds)
		}
	}

	loadRoutes(controls, *api)

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api.GET("/status", func(context *gin.Context) {
		context.Status(http.StatusAccepted)
	})

	gin.SetMode(gin.ReleaseMode)

	api.Use(gin.Recovery())

	return server
}
