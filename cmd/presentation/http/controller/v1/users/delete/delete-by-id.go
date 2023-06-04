package v1_delete_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/http/controller"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"
	"go-clean-api/cmd/presentation/http/middlewares"
)

type (
	deleteController struct {
		container *container.Container
	}
)

func New(c *container.Container) controller.Controller {
	return &deleteController{c}
}

func (deleteController *deleteController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "delete",
		Path:        "/:id",
		Middlewares: []controller.Middleware{middlewares.Log},
	}
}

func (createController *deleteController) Handle(req controller.HttpRequest) (*controller.HttpResponse, error) {
	id := req.Params.Get("id")

	err := createController.container.DeleteUserUseCase.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	return &controller.HttpResponse{
		Status: 204,
	}, nil
}

func (createController *deleteController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controller.HttpResponseError {
	if appErr != nil {
		if appErr.Code == domainexceptions.UserNotFound().Code {
			return httpexceptions.NotFound(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	return nil
}
