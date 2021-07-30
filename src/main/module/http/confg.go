package http_server

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	v1_user "go-api/src/presentation/http/controllers/v1/users"

	"github.com/gin-gonic/gin"
)

func loadControllers(container *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1_user.NewUserController(container),
	}
}

func loadRoutes(controls []controllers.Controller, api gin.RouterGroup) {
	for _, ctr := range controls {
		api_group := api.Group(ctr.PathGroup())

		for _, route := range ctr.LoadRoutes() {
			if len(route.Middlewares) > 0 {
				for _, mds := range route.Middlewares {
					middleware := func(gin_context *gin.Context) {
						mds(controllers.HttpRequest{
							Next: gin_context.Next,
						})
					}

					api_group.Use(middleware)
				}
			}

			function := func(gin_context *gin.Context) {
				var params controllers.Params

				if gin_context.Params != nil {
					for _, param := range gin_context.Params {
						param := controllers.Param{
							Key:   param.Key,
							Value: param.Value,
						}

						params = append(params, param)
					}
				}

				if route.Dto != nil {
					if err := gin_context.BindJSON(&route.Dto); err != nil {
						return
					}
				}

				result, err := route.Handle(controllers.HttpRequest{
					Body:    route.Dto,
					Params:  params,
					Query:   gin_context.Request.URL.Query(),
					Headers: gin_context.Request.Header,
				})

				if err.Data.Code != "" {
					gin_context.JSON(err.Status, err.Data)
				} else {
					if result.Headers != nil {
						for _, header := range result.Headers {
							gin_context.Header(header.Key, header.Value)
						}
					}

					if result.Data != nil {
						gin_context.JSON(result.Status, result.Data)
						return
					}

					gin_context.Status(result.Status)
				}
			}

			switch route.Method {
			case "get":
				api_group.GET(route.Path, function)
			case "post":
				api_group.POST(route.Path, function)
			case "put":
				api_group.PUT(route.Path, function)
			case "patch":
				api_group.PATCH(route.Path, function)
			case "delete":
				api_group.DELETE(route.Path, function)
			default:
			}
		}
	}
}
