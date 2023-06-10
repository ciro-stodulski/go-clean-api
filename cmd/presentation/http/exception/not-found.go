package httpexception

import "go-clean-api/cmd/presentation/http/controller"

func NotFound(data controller.HttpError) *controller.HttpResponseError {
	return &controller.HttpResponseError{
		Data:   data,
		Status: 404,
	}
}
