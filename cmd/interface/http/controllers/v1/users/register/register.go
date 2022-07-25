package controllerv1userregister

import (
	"go-api/cmd/core/entities/user"
	registeruserusecase "go-api/cmd/core/use-case/register-user"
	controllers "go-api/cmd/interface/http/controllers"
	ports_http "go-api/cmd/interface/http/ports"
	"log"

	"github.com/mitchellh/mapstructure"
)

func (rc *registerController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Handle:   rc.create,
		Dto:      registeruserusecase.Dto{},
	}
}

func (rc *registerController) create(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError) {
	dto := registeruserusecase.Dto{}
	mapstructure.Decode(req.Body, &dto)

	u, err := rc.container.RegisterUserUseCase.Register(dto)
	log.Default().Print(err)
	if err != nil {
		if err == user.ErrUserAlreadyExists {
			return nil, &ports_http.HttpResponseError{
				Data: ports_http.HttpError{
					Code:    "USER_ALREADY_EXISTS",
					Message: user.ErrUserAlreadyExists.Error(),
				},
				Status: 400,
			}
		}

		return nil, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}
	}

	log.Default().Print(u)

	return &ports_http.HttpResponse{
		Status: 201,
	}, nil
}
