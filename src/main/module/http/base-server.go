package http_server

import (
	controllers "go-api/src/presentation/http/controllers"
	ports_http "go-api/src/presentation/http/ports"

	"github.com/gin-gonic/gin"
)

func loadRoutes(controllers []controllers.Controller, api gin.RouterGroup) {
	for _, ctr := range controllers {
		route := ctr.LoadRoute()

		if route.PathRoot == "" {
			return
		}

		api_group := api.Group(route.PathRoot)

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

			if err != nil {
				status := 500

				if err.Status != 0 {
					status = err.Status
				}

				gin_context.JSON(status, err.Data)
			} else {
				if result.Headers != nil {
					for _, header := range result.Headers {
						gin_context.Header(header.Key, header.Value)
					}
				}

				status := 200

				if result.Status != 0 {
					status = result.Status
				}

				if result.Data != nil {
					gin_context.JSON(status, result.Data)
					return
				}

				gin_context.Status(status)
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
