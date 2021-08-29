package v1_user_create

import (
	"go-api/src/main/container"
	controllers "go-api/src/presentation/http/controllers"
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
)

type (
	createController struct {
		container *container.Container
	}
)

func NewController(c *container.Container) controllers.Controller {
	return &createController{c}
}

func (createController *createController) LoadRoute() controllers.CreateRoute {
	return controllers.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Handle:   createController.create,
		Dto:      create_dto.CreateDto{},
	}
}
