package v1_user

import (
	"go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/http/controller"
	httpexceptions "go-clean-api/cmd/presentation/http/exception"
	"go-clean-api/cmd/presentation/http/middlewares"
)

type (
	findByIdController struct {
		container *container.Container
	}
)

func New(c *container.Container) controller.Controller {
	return &findByIdController{c}
}

func (findByIdController *findByIdController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "get",
		Path:        "/:id",
		Middlewares: []controller.Middleware{middlewares.Log},
	}
}

func (findByIdController *findByIdController) Handle(req controller.HttpRequest) (*controller.HttpResponse, error) {
	id := req.Params.Get("id")

	u, err := findByIdController.container.GetUserUseCase.GetUser(id)

	if err != nil {
		return nil, err
	}

	return &controller.HttpResponse{
		Data:   u,
		Status: 200,
	}, nil
}

func (findByIdController *findByIdController) HandleError(appErr *exception.ApplicationException, err error) *controller.HttpResponseError {
	if appErr != nil {
		if appErr.Code == exception.UserNotFound().Code {
			return httpexceptions.NotFound(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	return nil
}
