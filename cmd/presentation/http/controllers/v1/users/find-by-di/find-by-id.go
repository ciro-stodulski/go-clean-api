package v1_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	"go-clean-api/cmd/presentation/http/middlewares"
	ports_http "go-clean-api/cmd/presentation/http/ports"
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
	if err == domainexceptions.ErrUserNotFound {
		return &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "USER_NOT_FOUND",
				Message: domainexceptions.ErrUserNotFound.Error(),
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
