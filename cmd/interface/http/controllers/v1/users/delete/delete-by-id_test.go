package v1_delete_user

import (
	"go-api/cmd/core/entities/user"
	ports_http "go-api/cmd/interface/http/ports"
	"go-api/cmd/main/container"
	deleteeuserusecasemock "go-api/cmd/shared/mocks/core/use-cases/delete-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_Delete(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockRepo := new(deleteeuserusecasemock.MockUseCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockRepo.On("DeleteUser", id).Return(nil)
		//

		// test func
		testService := New(&container.Container{
			DeleteUserUseCase: mockRepo,
		})
		result, err := testService.Handle(ports_http.HttpRequest{
			Params: ports_http.Params{
				ports_http.Param{Key: "id", Value: id},
			},
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &ports_http.HttpResponse{
			Status: 204,
		}, result)
		//
	})

	t.Run("error USER_NOT_FOUND", func(t *testing.T) {
		// make mock
		mockRepo := new(deleteeuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(&container.Container{
			DeleteUserUseCase: mockRepo,
		})

		err_http := testService.HandleError(user.ErrUserNotFound)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "USER_NOT_FOUND",
				Message: user.ErrUserNotFound.Error(),
			},
			Status: 404,
		}, err_http)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockRepo := new(deleteeuserusecasemock.MockUseCase)

		testService := New(&container.Container{
			DeleteUserUseCase: mockRepo,
		})
		//

		// test func
		err_http := testService.HandleError(user.ErrUserAlreadyExists)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}, err_http)
		//
	})
}
