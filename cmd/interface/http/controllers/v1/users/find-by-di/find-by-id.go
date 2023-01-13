package v1_user

import (
	entity_user "go-clean-api/cmd/core/entities/user"
	controllers "go-clean-api/cmd/interface/http/controllers"
	"go-clean-api/cmd/interface/http/middlewares"
	ports_http "go-clean-api/cmd/interface/http/ports"
	"go-clean-api/cmd/main/container"
)

type (
	findByIdController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &findByIdController{c}
}

func (findByIdController *findByIdController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "get",
		Path:        "/:id",
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}

func (findByIdController *findByIdController) Handle(req ports_http.HttpRequest) (*ports_http.HttpResponse, error) {

	findByIdController.container.ListUsersUseCase.ListUsers()

	return &ports_http.HttpResponse{
		Data:   "",
		Status: 200,
	}, nil
}

func (findByIdController *findByIdController) HandleError(err error) *ports_http.HttpResponseError {
	if err == entity_user.ErrUserNotFound {
		return &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "USER_NOT_FOUND",
				Message: entity_user.ErrUserNotFound.Error(),
			},
			Status: 404,
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
