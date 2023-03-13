package controllers

type (
	Controller interface {
		LoadRoute() CreateRoute
		Handle(req HttpRequest) (*HttpResponse, error)
		HandleError(err error) *HttpResponseError
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
