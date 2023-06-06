package controllerv1userregister

import (
	"errors"
	domaindto "go-clean-api/cmd/domain/dto"
	exception "go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/presentation/http/controller"
	httpexceptions "go-clean-api/cmd/presentation/http/exception"

	"go-clean-api/cmd/domain/entity/user"
	"go-clean-api/cmd/main/container"
	createuserusecasemock "go-clean-api/cmd/shared/mocks/application/use-cases/create-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_User_Register(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUse := new(createuserusecasemock.MockUseCase)

		dto := domaindto.Dto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockUse.On("Register", dto).Return(&user.User{}, nil)
		//

		// test func
		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})

		result, err := testService.Handle(controller.HttpRequest{
			Body: dto,
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &controller.HttpResponse{
			Status: 201,
		}, result)
		//
	})

	t.Run("error INVALID_DATA", func(t *testing.T) {
		// make mock
		mockUse := new(createuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})

		err_http := testService.HandleError(exception.InvalidEntity(), nil)
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

		mockUse := new(createuserusecasemock.MockUseCase)
		// test func

		//
		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})
		err_http := testService.HandleError(exception.UserAlreadyExists(), nil)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexceptions.Conflict(controller.HttpError{
			Code:    exception.UserAlreadyExists().Code,
			Message: exception.UserAlreadyExists().Message,
		}), err_http)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockUse := new(createuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})
		err_http := testService.HandleError(nil, errors.New("test"))
		//

		// asserts
		assert.Nil(t, err_http)
		//
	})
}
