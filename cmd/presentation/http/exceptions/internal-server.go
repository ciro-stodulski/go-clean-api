package httpexceptions

import "go-clean-api/cmd/presentation/http/controllers"

func InternalServer(data controllers.HttpError) *controllers.HttpResponseError {
	return &controllers.HttpResponseError{
		Data:   data,
		Status: 500,
	}
}
