package v1_delete_user

import (
	"errors"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"
	deleteeuserusecasemock "go-clean-api/cmd/shared/mocks/application/use-cases/delete-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_Delete(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUseCase := new(deleteeuserusecasemock.MockUseCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUseCase.On("DeleteUser", id).Return((*domainexceptions.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(&container.Container{
			DeleteUserUseCase: mockUseCase,
		})
		result, errApp, err := testService.Handle(controllers.HttpRequest{
			Params: controllers.Params{
				controllers.Param{Key: "id", Value: id},
			},
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.Nil(t, errApp)
		assert.NotNil(t, result)
		assert.Equal(t, &controllers.HttpResponse{
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

		errHttp := testService.HandleError(domainexceptions.UserNotFound(), nil)
		//

		// asserts
		assert.NotNil(t, errHttp)
		assert.Equal(t, httpexceptions.NotFound(controllers.HttpError{
			Code:    domainexceptions.UserNotFound().Code,
			Message: domainexceptions.UserNotFound().Message,
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
