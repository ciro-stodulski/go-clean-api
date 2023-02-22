package v1_delete_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presetation/http/controllers"
	"go-clean-api/cmd/presetation/http/middlewares"
	ports_http "go-clean-api/cmd/presetation/http/ports"
)

type (
	deleteController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &deleteController{c}
}

func (deleteController *deleteController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "delete",
		Path:        "/:id",
		Middlewares: []controllers.Middleware{middlewares.Log},
	}
}

func (createController *deleteController) Handle(req ports_http.HttpRequest) (*ports_http.HttpResponse, error) {
	id := req.Params.Get("id")

	err := createController.container.DeleteUserUseCase.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	return &ports_http.HttpResponse{
		Status: 204,
	}, nil
}

func (createController *deleteController) HandleError(err error) *ports_http.HttpResponseError {
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
