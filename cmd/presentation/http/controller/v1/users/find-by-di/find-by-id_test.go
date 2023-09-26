package findbyiduser

import (
	"go-clean-api/cmd/domain/entity/user"
	exception "go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/presentation/http/controller"
	httpexception "go-clean-api/cmd/presentation/http/exception"
	"go-clean-api/cmd/shared/mocks"
	usecasemock "go-clean-api/cmd/shared/mocks/application/use-case/use-case"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_User_Find_By_Id(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		userMock := mocks.NewMockUser()
		mockUse := new(usecasemock.MockUseCase[string, *user.User])
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUse.On("Perform", id).Return(userMock, nil)
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
		mockUse.AssertCalled(t, "Perform", id)
		assert.Equal(t, &controller.HttpResponse{
			Data:   userMock,
			Status: 200,
		}, result)
		//
	})

	t.Run("error USER_NOT_FOUND", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[string, *user.User])
		//

		// test func
		testService := New(mockUse)
		err_http := testService.HandleError(exception.UserNotFound())
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexception.NotFound(controller.HttpError{
			Code:    exception.UserNotFound().Code,
			Message: exception.UserNotFound().Message,
		}), err_http)
		//
	})

}
