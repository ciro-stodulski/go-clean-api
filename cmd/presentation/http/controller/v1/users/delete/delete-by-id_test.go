package v1_delete_user

import (
	"errors"
	exception "go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/presentation/http/controller"
	httpexception "go-clean-api/cmd/presentation/http/exception"
	deleteeuserusecasemock "go-clean-api/cmd/shared/mocks/application/use-case/delete-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_Delete(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUseCase := new(deleteeuserusecasemock.MockUseCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUseCase.On("DeleteUser", id).Return((*exception.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(&container.Container{
			DeleteUserUseCase: mockUseCase,
		})
		result, err := testService.Handle(controller.HttpRequest{
			Params: controller.Params{
				controller.Param{Key: "id", Value: id},
			},
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &controller.HttpResponse{
			Status: 204,
		}, result)
		//
	})

	t.Run("error USER_NOT_FOUND", func(t *testing.T) {
		// make mock
		mockUseCase := new(deleteeuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(&container.Container{
			DeleteUserUseCase: mockUseCase,
		})

		errHttp := testService.HandleError(exception.UserNotFound(), nil)
		//

		// asserts
		assert.NotNil(t, errHttp)
		assert.Equal(t, httpexception.NotFound(controller.HttpError{
			Code:    exception.UserNotFound().Code,
			Message: exception.UserNotFound().Message,
		}), errHttp)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockUseCase := new(deleteeuserusecasemock.MockUseCase)

		testService := New(&container.Container{
			DeleteUserUseCase: mockUseCase,
		})
		//

		// test func
		err_http := testService.HandleError(nil, errors.New("internal error"))
		//

		// asserts
		assert.Nil(t, err_http)
		//
	})
}
