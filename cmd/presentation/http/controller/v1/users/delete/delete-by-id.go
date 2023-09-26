package deleteuser

import (
	"go-clean-api/cmd/domain/exception"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/http/controller"
	httpexception "go-clean-api/cmd/presentation/http/exception"
	"go-clean-api/cmd/presentation/http/middlewares"
	"log"
)

type (
	deleteController struct {
		deleteUserUseCase usecase.UseCase[string, interface{}]
	}
)

func New(deleteUserUseCase usecase.UseCase[string, interface{}]) controller.Controller {
	return &deleteController{deleteUserUseCase}
}

// DeleteUser godoc
//
//	@Summary		Delete User
//	@Description	Delete a specific user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		204
//	@Failure		400	{object}	exception.ApplicationException	"{ "code": "USER_NOT_FOUND", "message": "User	not	found" }"
//	@Router			/v1/users/{id}   [delete]
func (deleteController *deleteController) LoadRoute() controller.CreateRoute {
	return controller.CreateRoute{
		PathRoot:    "/v1/users",
		Method:      "delete",
		Path:        "/:id",
		Middlewares: []controller.Middleware{middlewares.Log},
	}
}

func (createController *deleteController) Handle(req controller.HttpRequest) (*controller.HttpResponse, error) {
	id := req.Params.Get("id")

	_, err := createController.deleteUserUseCase.Perform(id)

	if err != nil {
		return nil, err
	}

	return &controller.HttpResponse{
		Status: 204,
	}, nil
}

func (createController *deleteController) HandleError(appErr *exception.ApplicationException) *controller.HttpResponseError {
	if appErr != nil {
		if appErr.Code == exception.UserNotFound().Code {
			return httpexception.NotFound(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	log.Default().Println("ERROR", appErr)

	return nil
}
