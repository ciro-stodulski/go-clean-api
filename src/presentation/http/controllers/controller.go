package controllers

import ports_http "go-api/src/presentation/http/controllers/ports"

type Controller interface {
	LoadRoutes() []CreateRoute
	PathGroup() string
}

type Middleware func(req ports_http.HttpRequest)

type CreateRoute struct {
	Method      string
	Path        string
	Middlewares []Middleware
	Dto         interface{}
	Handle      func(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError)
}
