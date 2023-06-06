package httpexceptions

import "go-clean-api/cmd/presentation/http/controller"

func Conflict(data controller.HttpError) *controller.HttpResponseError {
	return &controller.HttpResponseError{
		Data:   data,
		Status: 409,
	}
}
