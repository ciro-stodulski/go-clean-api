package v1_user_create

import (
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
	ports_http "go-api/src/presentation/http/ports"

	"github.com/mitchellh/mapstructure"
)

func (createController *createController) create(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError) {
	dto := create_dto.CreateDto{}
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
