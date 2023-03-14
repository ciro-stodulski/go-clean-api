package v1_delete_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"
	"go-clean-api/cmd/presentation/http/middlewares"
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

func (createController *deleteController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, *domainexceptions.ApplicationException, error) {
	id := req.Params.Get("id")

	errApp, err := createController.container.DeleteUserUseCase.DeleteUser(id)

	if errApp != nil || err != nil {
		return nil, errApp, err
	}

	return &controllers.HttpResponse{
		Status: 204,
	}, nil, nil
}

func (createController *deleteController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controllers.HttpResponseError {
	if appErr != nil {
		if appErr.Code == domainexceptions.UserNotFound().Code {
			return httpexceptions.NotFound(controllers.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	return nil
}
