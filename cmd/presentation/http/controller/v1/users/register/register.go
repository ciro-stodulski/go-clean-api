package controllerv1userregister

import (
	"go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"
	"go-clean-api/cmd/domain/exception"
	usecase "go-clean-api/cmd/domain/use-case"

	"go-clean-api/cmd/presentation/http/controller"
	httpexceptions "go-clean-api/cmd/presentation/http/exception"

	"log"
)

type (
	registerController struct {
		registerUserUseCase usecase.UseCase[dto.RegisterUser, *user.User]
	}
)

func New(registerUserUseCase usecase.UseCase[dto.RegisterUser, *user.User]) controller.Controller {
	return &registerController{registerUserUseCase}
}

// Register User godoc
//
//	@Summary		Register User
//	@Description	Create a new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body	dto.RegisterUser	true	"User Data"
//	@Success		201
//	@Failure		400	{object}	exception.ApplicationException	"{ "code": "INVALID_ENTITY", "message": "Invalid	entity" },{ "code": "USER_ALREADY_EXISTS", "message": "Already	exists	user" }"
//	@Router			/v1/users [post]
func (rc *registerController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot: "/v1/users",
		Method:   "post",
		Path:     "/",
		Dto:      &dto.RegisterUser{},
	}
}

func (rc *registerController) Handle(req controller.HttpRequest) (*controller.HttpResponse[any], error) {
	_, err := rc.registerUserUseCase.Perform(req.Body.(dto.RegisterUser))

	if err != nil {
		return nil, err
	}

	return &controller.HttpResponse[any]{
		Status: 201,
	}, nil
}

func (rc *registerController) HandleError(appErr *exception.ApplicationException) *controller.HttpResponse[controller.HttpError] {
	if appErr != nil {
		if appErr.Code == exception.InvalidEntity().Code {
			return httpexceptions.BadRequest(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}

		if appErr.Code == exception.UserAlreadyExists().Code {
			return httpexceptions.Conflict(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	log.Default().Println("internal error", appErr)

	return nil
}
