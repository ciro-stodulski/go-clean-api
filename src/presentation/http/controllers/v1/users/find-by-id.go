package v1_user

import (
	entity_user "go-api/src/core/entities/user"
	ports_http "go-api/src/presentation/http/ports"
)

func (createController *createController) findById(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError) {
	id := req.Params.Get("id")

	user, err := createController.container.UserService.GetUser(id)

	if err != nil {
		if err == entity_user.ErrUserNotFound {
			return nil, &ports_http.HttpResponseError{
				Data: ports_http.HttpError{
					Code:    "USER_NOT_FOUND",
					Message: entity_user.ErrUserNotFound.Error(),
				},
				Status: 400,
			}
		}

		return nil, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}
	}

	return &ports_http.HttpResponse{
		Data:   user,
		Status: 200,
	}, nil
}
