package http_server

import (
	ports_http "go-api/src/interface/http/ports"
	"go-api/src/main/container"
	"go-api/src/shared/env"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Engine *gin.Engine
}

func (server *HttpServer) Start() error {
	log.Default().Print("Http: Server started with succeffully")

	return server.Engine.Run(env.Env().HostHttp + ":" + env.Env().HostPort)
}

func (server *HttpServer) New(container *container.Container) {
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
}
