package http

import (
	"go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/presentation/http/controller"
	"go-clean-api/docs"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func loadRoutes(controllers []controller.Controller, api gin.RouterGroup) {
	for _, ctr := range controllers {
		route := ctr
		route_config := ctr.LoadRoute()

		if route_config.PathRoot == "" {
			return
		}

		api_group := api.Group(route_config.PathRoot)
		docs.SwaggerInfo.BasePath = route_config.PathRoot

		loadMiddlewares(route_config, api_group)

		var dto reflect.Value

		function := func(gin_context *gin.Context) {

			params := loadParams(gin_context)

			if route_config.Dto != nil {

				dtoType := reflect.TypeOf(route_config.Dto).Elem()

				dto = reflect.New(dtoType).Elem()

				if route_config.Dto != nil {
					gin_context.Request.FormFile("*")
					if gin_context.Request.MultipartForm == nil {
						if err := gin_context.ShouldBindJSON(dto.Addr().Interface()); err != nil {
							handleValidationErrors(gin_context, err)
							return
						}
					} else {
						if err := gin_context.ShouldBind(dto.Addr().Interface()); err != nil {
							handleValidationErrors(gin_context, err)
							return
						}
					}
				}
			}

			var bodyValue interface{}
			if route_config.Dto != nil {

				bodyValue = dto.Interface()
			} else {
				bodyValue = nil
			}

			result, err := route.Handle(controller.HttpRequest{
				Body:    bodyValue,
				Params:  params,
				Query:   gin_context.Request.URL.Query(),
				Headers: gin_context.Request.Header,
			})

			if err != nil {
				var result_error *controller.HttpResponseError

				if appErr, ok := err.(*exception.ApplicationException); ok {
					result_error = route.HandleError(appErr)
				} else {
					log.Default().Println("INTERNAL_SERVER_ERROR", err)

					result_error = &controller.HttpResponseError{
						Data:   controller.HttpError{Code: "INTERNAL_SERVER_ERROR", Message: "internal server error"},
						Status: 500,
					}
				}

				gin_context.JSON(result_error.Status, result_error.Data)
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

		switch route_config.Method {
		case "get":
			api_group.GET(route_config.Path, function)
		case "post":
			api_group.POST(route_config.Path, function)
		case "put":
			api_group.PUT(route_config.Path, function)
		case "patch":
			api_group.PATCH(route_config.Path, function)
		case "delete":
			api_group.DELETE(route_config.Path, function)
		default:

		}

	}

}

func loadParams(context *gin.Context) controller.Params {
	var params controller.Params

	if context.Params != nil {
		for _, param := range context.Params {
			param := controller.Param{
				Key:   param.Key,
				Value: param.Value,
			}

			params = append(params, param)
		}
	}

	return params
}

func loadMiddlewares(route controller.CreateRoute, api_group *gin.RouterGroup) {
	if len(route.Middlewares) > 0 {

		for _, mds := range route.Middlewares {
			middleware := func(gin_context *gin.Context) {
				params := loadParams(gin_context)

				if route.Dto != nil {
					if err := gin_context.BindJSON(&route.Dto); err != nil {
						return
					}
				}

				mds(controller.HttpRequest{
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

func handleValidationErrors(c *gin.Context, err error) {
	var errorDetails []map[string]interface{}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusBadRequest, controller.HttpError{
			Code:    "INVALID_SCHEMA",
			Message: "Invalid schema payload",
		})
		return
	}

	for _, validationError := range validationErrors {
		fieldName := validationError.Field()
		errorMessage := validationError.Error()

		fieldDetails := map[string]interface{}{
			"campo": fieldName,
			"erro":  errorMessage,
		}

		errorDetails = append(errorDetails, fieldDetails)
	}

	c.JSON(http.StatusBadRequest, controller.HttpError{
		Code:    "INVALID_SCHEMA",
		Message: "Invalid schema payload",
		Detail:  errorDetails,
	})
}
