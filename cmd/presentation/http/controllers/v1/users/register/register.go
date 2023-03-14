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

func (rc *registerController) Handle(req controllers.HttpRequest) (*controllers.HttpResponse, *domainexceptions.ApplicationException, error) {
	dto := domaindto.Dto{}
	mapstructure.Decode(req.Body, &dto)

	_, errApp, err := rc.container.RegisterUserUseCase.Register(dto)

	if err != nil || errApp != nil {
		return nil, errApp, err
	}

	return &controllers.HttpResponse{
		Status: 201,
	}, nil, nil
}

func (rc *registerController) HandleError(appErr *domainexceptions.ApplicationException, err error) *controllers.HttpResponseError {
	if appErr != nil {
		if appErr.Code == domainexceptions.InvalidEntity().Code {
			return httpexceptions.BadRequest(controllers.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}

		if appErr.Code == domainexceptions.UserAlreadyExists().Code {
			return httpexceptions.Conflict(controllers.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	log.Default().Println(err)

	return nil
}
