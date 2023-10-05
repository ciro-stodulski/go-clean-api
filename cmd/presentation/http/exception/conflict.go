package httpexception

import "go-clean-api/cmd/presentation/http/controller"

func Conflict(data controller.HttpError) *controller.HttpResponse[controller.HttpError] {
	return &controller.HttpResponse[controller.HttpError]{
		Data:   data,
		Status: 409,
	}
}
