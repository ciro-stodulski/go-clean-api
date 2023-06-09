package findbyiduser

import (
	"errors"
	exception "go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/presentation/http/controller"
	httpexception "go-clean-api/cmd/presentation/http/exception"
	"go-clean-api/cmd/shared/mocks"
	getuserusecasemock "go-clean-api/cmd/shared/mocks/application/use-case/get-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_User_Find_By_Id(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		userMock := mocks.NewMockUser()
		mockUse := new(getuserusecasemock.MockUseCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUse.On("GetUser", id).Return(userMock, nil)
		//

		// test func
		testService := New(mockUse)

		result, err := testService.Handle(controller.HttpRequest{
			Params: controller.Params{
				controller.Param{Key: "id", Value: id},
			},
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		mockUse.AssertCalled(t, "GetUser", id)
		assert.Equal(t, &controller.HttpResponse{
			Data:   userMock,
			Status: 200,
		}, result)
		//
	})

	t.Run("error USER_NOT_FOUND", func(t *testing.T) {
		// make mock
		mockUse := new(getuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(mockUse)
		err_http := testService.HandleError(exception.UserNotFound(), nil)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexception.NotFound(controller.HttpError{
			Code:    exception.UserNotFound().Code,
			Message: exception.UserNotFound().Message,
		}), err_http)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockUse := new(getuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(mockUse)

		err_http := testService.HandleError(nil, errors.New("internal error"))
		//

		// asserts
		assert.Nil(t, err_http)
		//
	})
}