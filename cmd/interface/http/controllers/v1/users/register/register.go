package controllerv1userregister

import (
	entity "go-clean-api/cmd/core/entities"
	"go-clean-api/cmd/core/entities/user"
	registeruserusecase "go-clean-api/cmd/core/use-case/register-user"
	controllers "go-clean-api/cmd/interface/http/controllers"
	ports_http "go-clean-api/cmd/interface/http/ports"
	"go-clean-api/cmd/main/container"
	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	registerController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &registerController{c}
}

func (rc *registerController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Dto:      registeruserusecase.Dto{},
	}
}

func (rc *registerController) Handle(req ports_http.HttpRequest) (*ports_http.HttpResponse, error) {
	dto := registeruserusecase.Dto{}
	mapstructure.Decode(req.Body, &dto)

	_, err := rc.container.RegisterUserUseCase.Register(dto)

	if err != nil {
		log.Default().Print(err)

		return nil, err
	}

	return &ports_http.HttpResponse{
		Status: 201,
	}, nil
}

func (rc *registerController) HandleError(err error) *ports_http.HttpResponseError {
	if err == entity.ErrInvalidEntity {
		return &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INVALID_DATA",
				Message: entity.ErrInvalidEntity.Error(),
			},
			Status: 400,
		}
	}

	if err == user.ErrUserAlreadyExists {
		return &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "USER_ALREADY_EXISTS",
				Message: user.ErrUserAlreadyExists.Error(),
			},
			Status: 400,
		}
	}

	return &ports_http.HttpResponseError{
		Data: ports_http.HttpError{
			Code:    "INTERNAL_ERROR",
			Message: "internal error",
		},
		Status: 500,
	}
}
