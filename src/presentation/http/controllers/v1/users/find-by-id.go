package v1_user

import (
	entity_user "go-api/src/core/entities/user"
	controllers "go-api/src/presentation/http/controllers"
)

func (createController *createController) findById(req controllers.HttpRequest) (controllers.HttpResponse, controllers.HttpResponseError) {
	id := req.Params.Get("id")

	user, err := createController.container.UserService.GetUser(id)

	if err != nil {
		if err == entity_user.ErrUserNotFound {
			return controllers.HttpResponse{}, controllers.HttpResponseError{
				Data: controllers.HttpError{
					Code:    "USER_NOT_FOUND",
					Message: entity_user.ErrUserNotFound.Error(),
				},
				Status: 400,
			}
		}

		return controllers.HttpResponse{}, controllers.HttpResponseError{
			Data: controllers.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}
	}

	return controllers.HttpResponse{
		Data:   user,
		Status: 200,
	}, controllers.HttpResponseError{}
}
