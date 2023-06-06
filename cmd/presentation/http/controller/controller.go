package controller

import "go-clean-api/cmd/domain/exception"

type (
	Controller interface {
		LoadRoute() CreateRoute
		Handle(req HttpRequest) (*HttpResponse, error)
		HandleError(appErr *exception.ApplicationException, err error) *HttpResponseError
	}

	Middleware func(req HttpRequest)

	CreateRoute struct {
		PathRoot    string
		Method      string
		Path        string
		Middlewares []Middleware
		Dto         interface{}
	}
)
