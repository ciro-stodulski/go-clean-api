package http

import (
	ports_http "go-clean-api/cmd/interface/http/ports"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/main/modules"
	"go-clean-api/cmd/shared/env"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
				middleware(ports_http.HttpRequest{
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

	api.GET("/status", func(context *gin.Context) {
		context.Status(http.StatusAccepted)
	})

	gin.SetMode(gin.ReleaseMode)

	api.Use(gin.Recovery())

	return server
}
