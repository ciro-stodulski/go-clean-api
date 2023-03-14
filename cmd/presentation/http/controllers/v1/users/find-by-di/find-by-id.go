package v1_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"
	"go-clean-api/cmd/presentation/http/middlewares"
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

func (findByIdController *findByIdController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, *domainexceptions.ApplicationException, error) {
	id := req.Params.Get("id")

	u, errApp, err := findByIdController.container.GetUserUseCase.GetUser(id)

	if err != nil || errApp != nil {
		return nil, errApp, err
	}

	return &controllers.HttpResponse{
		Data:   u,
		Status: 200,
	}, nil, nil
}

func (findByIdController *findByIdController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controllers.HttpResponseError {
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
