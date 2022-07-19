package v1_delete_user

import (
	entity_user "go-api/src/core/entities/user"
	ports_http "go-api/src/interface/http/ports"
)

func (createController *deleteController) deleteById(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError) {
	id := req.Params.Get("id")

	err := createController.container.DeleteUserUseCase.DeleteUser(id)
	if err != nil {
		if err == entity_user.ErrUserNotFound {
			return nil, &ports_http.HttpResponseError{
				Data: ports_http.HttpError{
					Code:    "USER_NOT_FOUND",
					Message: entity_user.ErrUserNotFound.Error(),
				},
				Status: 404,
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
		Status: 204,
	}, nil
}
