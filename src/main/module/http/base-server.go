package http_server

import (
	controllers "go-api/src/presentation/http/controllers"
	ports_http "go-api/src/presentation/http/controllers/ports"

	"github.com/gin-gonic/gin"
)

func loadRoutes(controls []controllers.Controller, api gin.RouterGroup) {
	for _, ctr := range controls {
		api_group := api.Group(ctr.PathGroup())

		for _, route := range ctr.LoadRoutes() {
			loadMiddlewares(route, api_group)

			function := func(gin_context *gin.Context) {
				params := loadParams(gin_context)

				if route.Dto != nil {
					if err := gin_context.BindJSON(&route.Dto); err != nil {
						return
					}
				}

				result, err := route.Handle(ports_http.HttpRequest{
					Body:    route.Dto,
					Params:  params,
					Query:   gin_context.Request.URL.Query(),
					Headers: gin_context.Request.Header,
				})

				if err.Data != (ports_http.HttpError{}) {
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

func loadParams(context *gin.Context) ports_http.Params {
	var params ports_http.Params

	if context.Params != nil {
		for _, param := range context.Params {
			param := ports_http.Param{
				Key:   param.Key,
				Value: param.Value,
			}

			params = append(params, param)
		}
	}

	return params
}

func loadMiddlewares(route controllers.CreateRoute, api_group *gin.RouterGroup) {
	if len(route.Middlewares) > 0 {

		for _, mds := range route.Middlewares {
			middleware := func(gin_context *gin.Context) {
				params := loadParams(gin_context)

				if route.Dto != nil {
					if err := gin_context.BindJSON(&route.Dto); err != nil {
						return
					}
				}

				mds(ports_http.HttpRequest{
					Params:  params,
					Query:   gin_context.Request.URL.Query(),
					Headers: gin_context.Request.Header,
					Next:    gin_context.Next,
					Body:    route.Dto,
				})
			}

			api_group.Use(middleware)
		}
	}
}
