package http

import (
	"encoding/json"
	"fmt"
	"go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/presentation/http/controller"
	"go-clean-api/docs"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func loadRoutes(controllers []controller.Controller, api gin.RouterGroup) {
	hub := NewSSEHub()

	for _, ctr := range controllers {
		route := ctr
		routeConfig := ctr.LoadRoute()

		if routeConfig.PathRoot == "" {
			return
		}

		api_group := api.Group(routeConfig.PathRoot)
		docs.SwaggerInfo.BasePath = routeConfig.PathRoot

		loadMiddlewares(routeConfig, api_group)

		var dto reflect.Value

		function := func(c *gin.Context) {

			params := loadParams(c)

			if routeConfig.Dto != nil {

				dtoType := reflect.TypeOf(routeConfig.Dto).Elem()

				dto = reflect.New(dtoType).Elem()

				if routeConfig.Dto != nil {
					c.Request.FormFile("*")
					if c.Request.MultipartForm == nil {
						if err := c.ShouldBindJSON(dto.Addr().Interface()); err != nil {
							handleValidationErrors(c, err)
							return
						}
					} else {
						if err := c.ShouldBind(dto.Addr().Interface()); err != nil {
							handleValidationErrors(c, err)
							return
						}
					}
				}
			}

			var bodyValue any
			if routeConfig.Dto != nil {

				bodyValue = dto.Interface()
			} else {
				bodyValue = nil
			}

			var result *controller.HttpResponse[any]
			var err error

			if route.LoadRoute().IsServerSentEvents {
				serverSentEvents(c, route, hub, bodyValue, params)
			} else {
				result, err = route.Handle(controller.HttpRequest{
					Body:    bodyValue,
					Params:  params,
					Query:   c.Request.URL.Query(),
					Headers: c.Request.Header,
				})
			}

			if err != nil {
				var result_error *controller.HttpResponse[controller.HttpError]

				if appErr, ok := err.(*exception.ApplicationException); ok {
					result_error = route.HandleError(appErr)
				}

				if result_error == nil {
					log.Default().Println("INTERNAL_SERVER_ERROR", err)

					result_error = &controller.HttpResponse[controller.HttpError]{
						Data:   controller.HttpError{Code: "INTERNAL_SERVER_ERROR", Message: "internal server error"},
						Status: 500,
					}
				}

				c.JSON(result_error.Status, result_error.Data)
			} else {
				if result.Headers != nil {
					for _, header := range result.Headers {
						c.Header(header.Key, header.Value)
					}
				}

				status := 200

				if result.Status != 0 {
					status = result.Status
				}

				if result.Data != nil {
					c.JSON(status, result.Data)
					return
				}
				c.Status(status)
			}

		}

		switch routeConfig.Method {
		case "get":
			api_group.GET(routeConfig.Path, function)
		case "post":
			api_group.POST(routeConfig.Path, function)
		case "put":
			api_group.PUT(routeConfig.Path, function)
		case "patch":
			api_group.PATCH(routeConfig.Path, function)
		case "delete":
			api_group.DELETE(routeConfig.Path, function)
		default:

		}

	}

}

func serverSentEvents(c *gin.Context, route controller.Controller, hub *SSEHub, bodyValue any, params controller.Params) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	messageChannel := make(chan any)

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		err := &controller.HttpResponse[controller.HttpError]{
			Data:   controller.HttpError{Code: "INTERNAL_SERVER_ERROR", Message: "Stream error"},
			Status: 500,
		}
		c.JSON(err.Status, err.Data)
		return
	}

	notify := c.Writer.CloseNotify()
	go func() {
		log.Default().Printf("[HTTP MODULE]{ServerSentEvents} close connection %v", <-notify)
		defer close(messageChannel)

		hub.RemoveClient(c)
		<-notify
	}()

	hub.AddClient(c, messageChannel)

	go func() {
		log.Default().Printf("[HTTP MODULE]{ServerSentEvents} open connection")

		for {
			result, err := route.Handle(controller.HttpRequest{
				Body:    bodyValue,
				Params:  params,
				Query:   c.Request.URL.Query(),
				Headers: c.Request.Header,
			})

			if err != nil {
				hub.Broadcast(err)

			} else {
				response, err := json.Marshal(result.Data)
				if err != nil {
					panic(err)
				}
				hub.Broadcast(response)

			}

			time.Sleep(time.Duration(route.LoadRoute().TimeSecondsSentEvents) * time.Second)
		}
	}()

	// Envie mensagens SSE para o cliente
	for message := range messageChannel {
		fmt.Fprintf(c.Writer, "data: %s\n\n", message)
		flusher.Flush()
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
			middleware := func(c *gin.Context) {
				params := loadParams(c)

				if route.Dto != nil {
					if err := c.BindJSON(&route.Dto); err != nil {
						return
					}
				}

				mds(controller.HttpRequest{
					Params:  params,
					Query:   c.Request.URL.Query(),
					Headers: c.Request.Header,
					Next:    c.Next,
					Body:    route.Dto,
				})
			}

			api_group.Use(middleware)
		}
	}
}

func handleValidationErrors(c *gin.Context, err error) {
	var errorDetails []map[string]any

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

		fieldDetails := map[string]any{
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
