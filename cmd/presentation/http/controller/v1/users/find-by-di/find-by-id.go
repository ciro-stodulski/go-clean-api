package findbyiduser

import (
	"go-clean-api/cmd/domain/entity/user"
	"go-clean-api/cmd/domain/exception"
	usecase "go-clean-api/cmd/domain/use-case"
	"go-clean-api/cmd/presentation/http/controller"
	httpexception "go-clean-api/cmd/presentation/http/exception"
	"go-clean-api/cmd/presentation/http/middlewares"
	"log"
)

type (
	findByIdController struct {
		getUserUseCase usecase.IUseCase[string, *user.User]
	}
)

func New(getUserUseCase usecase.IUseCase[string, *user.User]) controller.Controller {
	return &findByIdController{getUserUseCase}
}

// GetUserByID godoc
//
//	@Summary		Get User
//	@Description	Get information about a specific user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	user.User
//
//	@Failure		400	{object}	exception.ApplicationException	"{ "code": "USER_NOT_FOUND", "message": "User	not	found" }"
//
//	@Router			/v1/users/{id}   [get]
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

	u, err := findByIdController.getUserUseCase.Perform(id)

	if err != nil {
		return nil, err
	}

	return &controller.HttpResponse{
		Data:   u,
		Status: 200,
	}, nil
}

func (findByIdController *findByIdController) HandleError(appErr *exception.ApplicationException) *controller.HttpResponseError {
	if appErr != nil {
		if appErr.Code == exception.UserNotFound().Code {
			return httpexception.NotFound(controller.HttpError{
				Code:    appErr.Code,
				Message: appErr.Message,
			})
		}
	}

	log.Default().Println("internal error", appErr)

	return nil
}
