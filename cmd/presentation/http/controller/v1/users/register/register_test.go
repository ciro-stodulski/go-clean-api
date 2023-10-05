package controllerv1userregister

import (
	"go-clean-api/cmd/domain/dto"
	exception "go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/presentation/http/controller"
	httpexceptions "go-clean-api/cmd/presentation/http/exception"

	"go-clean-api/cmd/domain/entity/user"
	usecasemock "go-clean-api/cmd/shared/mocks/application/use-case/use-case"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_User_Register(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[dto.RegisterUser, *user.User])

		dto := dto.RegisterUser{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockUse.On("Perform", dto).Return(&user.User{}, nil)
		//

		// test func
		testService := New(mockUse)

		result, err := testService.Handle(controller.HttpRequest{
			Body: dto,
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &controller.HttpResponse[any]{
			Status: 201,
		}, result)
		//
	})

	t.Run("error INVALID_DATA", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[dto.RegisterUser, *user.User])
		//

		// test func
		testService := New(mockUse)

		err_http := testService.HandleError(exception.InvalidEntity())
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexceptions.BadRequest(controller.HttpError{
			Code:    exception.InvalidEntity().Code,
			Message: exception.InvalidEntity().Message,
		}), err_http)
		//
	})

	t.Run("error USER_ALREADY_EXISTS", func(t *testing.T) {
		// make mock

		mockUse := new(usecasemock.MockUseCase[dto.RegisterUser, *user.User])
		// test func

		//
		testService := New(mockUse)
		err_http := testService.HandleError(exception.UserAlreadyExists())
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexceptions.Conflict(controller.HttpError{
			Code:    exception.UserAlreadyExists().Code,
			Message: exception.UserAlreadyExists().Message,
		}), err_http)
		//
	})

}
