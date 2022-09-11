package controllers

import ports_http "go-clean-api/cmd/interface/http/ports"

type (
	Controller interface {
		LoadRoute() CreateRoute
		Handle(req ports_http.HttpRequest) (*ports_http.HttpResponse, error)
		HandleError(err error) *ports_http.HttpResponseError
	}

	Middleware func(req ports_http.HttpRequest)

	CreateRoute struct {
		PathRoot    string
		Method      string
		Path        string
		Middlewares []Middleware
		Dto         interface{}
	}
)
