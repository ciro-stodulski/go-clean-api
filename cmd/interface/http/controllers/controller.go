package controllers

import ports_http "go-api/cmd/interface/http/ports"

type (
	Controller interface {
		LoadRoute() CreateRoute
	}

	Middleware func(req ports_http.HttpRequest)

	CreateRoute struct {
		PathRoot    string
		Method      string
		Path        string
		Middlewares []Middleware
		Dto         interface{}
		Handle      func(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError)
	}
)
