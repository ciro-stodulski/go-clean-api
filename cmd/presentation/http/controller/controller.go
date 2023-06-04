package controller

import domainexceptions "go-clean-api/cmd/domain/exceptions"

type (
	Controller interface {
		LoadRoute() CreateRoute
		Handle(req HttpRequest) (*HttpResponse, error)
		HandleError(appErr *domainexceptions.ApplicationException, err error) *HttpResponseError
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
