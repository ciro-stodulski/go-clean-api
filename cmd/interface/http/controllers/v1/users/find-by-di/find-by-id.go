package v1_user

import (
	entity_user "go-api/cmd/core/entities/user"
	controllers "go-api/cmd/interface/http/controllers"
	"go-api/cmd/interface/http/middlewares"
	ports_http "go-api/cmd/interface/http/ports"
	"go-api/cmd/main/container"
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
	id := req.Params.Get("id")

	u, err := findByIdController.container.GetUserUseCase.GetUser(id)

	if err != nil {
		return nil, err
	}

	return &ports_http.HttpResponse{
		Data:   u,
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
