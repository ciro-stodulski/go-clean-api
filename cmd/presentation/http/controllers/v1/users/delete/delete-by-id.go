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

func (createController *deleteController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, error) {
	id := req.Params.Get("id")

	err := createController.container.DeleteUserUseCase.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	return &controllers.HttpResponse{
		Status: 204,
	}, nil
}

func (createController *deleteController) HandleError(err error) *controllers.HttpResponseError {
	if err == domainexceptions.ErrUserNotFound {
		return httpexceptions.NotFound(controllers.HttpError{
			Code:    "USER_NOT_FOUND",
			Message: domainexceptions.ErrUserNotFound.Error(),
		})
	}

	return httpexceptions.InternalServer(controllers.HttpError{
		Code:    "INTERNAL_ERROR",
		Message: "internal error",
	})
}
