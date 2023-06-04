package controllerv1userregister

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/http/controller"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"

	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	registerController struct {
		container *container.Container
	}
)

func New(c *container.Container) controller.Controller {
	return &registerController{c}
}

func (rc *registerController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Dto:      domaindto.Dto{},
	}
}

func (rc *registerController) Handle(req controller.HttpRequest) (*controller.HttpResponse, error) {
	dto := domaindto.Dto{}
	mapstructure.Decode(req.Body, &dto)

	_, err := rc.container.RegisterUserUseCase.Register(dto)

	if err != nil {
		return nil, err
	}

	return &controller.HttpResponse{
		Status: 201,
	}, nil
}

func (rc *registerController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controller.HttpResponseError {
	if appErr != nil {
		if appErr.Code == domainexceptions.InvalidEntity().Code {
			return httpexceptions.BadRequest(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}

		if appErr.Code == domainexceptions.UserAlreadyExists().Code {
			return httpexceptions.Conflict(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	log.Default().Println(err)

	return nil
}
