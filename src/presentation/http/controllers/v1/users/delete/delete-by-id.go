package v1_delete_user

import (
	ports_http "go-api/src/presentation/http/ports"
)

func (createController *createController) deleteById(req ports_http.HttpRequest) (*ports_http.HttpResponse, *ports_http.HttpResponseError) {
	id := req.Params.Get("id")

	err := createController.container.DeleteUserUseCase.DeleteUser(id)
	if err != nil {
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
