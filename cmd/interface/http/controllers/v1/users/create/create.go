package v1_user_create

import (
	"go-api/cmd/core/ports"
	ports_http "go-api/cmd/interface/http/ports"

	"github.com/mitchellh/mapstructure"
)

func (createController *createController) create(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError) {
	dto := ports.CreateDto{}
	mapstructure.Decode(req.Body, &dto)

	err := createController.container.CreateUserProducerUseCase.CreateUser(dto)

	if err != nil {
		return nil, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}
	}

	return &ports_http.HttpResponse{
		Status: 200,
	}, nil
}
