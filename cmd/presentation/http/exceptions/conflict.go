package httpexceptions

import "go-clean-api/cmd/presentation/http/controllers"

func Conflict(data controllers.HttpError) *controllers.HttpResponseError {
	return &controllers.HttpResponseError{
		Data:   data,
		Status: 409,
	}
}
