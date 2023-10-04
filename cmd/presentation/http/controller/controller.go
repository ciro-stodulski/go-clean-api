package controller

import "go-clean-api/cmd/domain/exception"

type (
	ChannelManager[Channel any] struct {
		IsOpen  bool
		Channel chan Channel
	}
	Controller interface {
		LoadRoute() CreateRoute
		Handle(req HttpRequest) (*HttpResponse, error)
		HandleError(appErr *exception.ApplicationException) *HttpResponseError
	}

	Middleware func(req HttpRequest)

	CreateRoute struct {
		IsServerSentEvents    bool
		TimeSecondsSentEvents int
		PathRoot              string
		Method                string
		Path                  string
		Middlewares           []Middleware
		Dto                   interface{}
	}
)
