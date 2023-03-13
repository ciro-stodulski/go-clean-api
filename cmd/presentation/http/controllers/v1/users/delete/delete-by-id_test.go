package v1_delete_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	deleteeuserusecasemock "go-clean-api/cmd/shared/mocks/application/use-cases/delete-user"
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
		result, err := testService.Handle(controllers.HttpRequest{
			Params: controllers.Params{
				controllers.Param{Key: "id", Value: id},
			},
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &controllers.HttpResponse{
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

		err_http := testService.HandleError(domainexceptions.ErrUserNotFound)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &controllers.HttpResponseError{
			Data: controllers.HttpError{
				Code:    "USER_NOT_FOUND",
				Message: domainexceptions.ErrUserNotFound.Error(),
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
		err_http := testService.HandleError(domainexceptions.ErrUserAlreadyExists)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &controllers.HttpResponseError{
			Data: controllers.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}, err_http)
		//
	})
}
