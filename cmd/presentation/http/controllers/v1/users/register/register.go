package controllerv1userregister

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"

	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	registerController struct {
		container *container.Container
	}
)

func New(c *container.Container) controllers.Controller {
	return &registerController{c}
}

func (rc *registerController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Dto:      domaindto.Dto{},
	}
}

func (rc *registerController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, error) {
	dto := domaindto.Dto{}
	mapstructure.Decode(req.Body, &dto)

	_, err := rc.container.RegisterUserUseCase.Register(dto)

	if err != nil {
		log.Default().Print(err)

		return nil, err
	}

	return &controllers.HttpResponse{
		Status: 201,
	}, nil
}

func (rc *registerController) HandleError(err error) *controllers.HttpResponseError {
	if err.Error() == domainexceptions.InvalidEntity().Error() {
		return httpexceptions.BadRequest(controllers.HttpError{
			Code:    "INVALID_DATA",
			Message: err.Error(),
		})
	}

	if err.Error() == domainexceptions.UserAlreadyExists().Error() {
		return httpexceptions.BadRequest(controllers.HttpError{
			Code:    "USER_ALREADY_EXISTS",
			Message: err.Error(),
		})
	}

	return httpexceptions.InternalServer(controllers.HttpError{
		Code:    "INTERNAL_ERROR",
		Message: "internal error",
	})
}
