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

func (findByIdController *findByIdController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, error) {
	id := req.Params.Get("id")

	u, err := findByIdController.container.GetUserUseCase.GetUser(id)

	if err != nil {
		return nil, err
	}

	return &controllers.HttpResponse{
		Data:   u,
		Status: 200,
	}, nil
}

func (findByIdController *findByIdController) HandleError(err error) *controllers.HttpResponseError {
	if err.Error() == domainexceptions.UserNotFound().Error() {
		return httpexceptions.NotFound(controllers.HttpError{
			Code:    "USER_NOT_FOUND",
			Message: domainexceptions.UserNotFound().Error(),
		})
	}

	return httpexceptions.InternalServer(controllers.HttpError{
		Code:    "INTERNAL_ERROR",
		Message: "internal error",
	})
}
