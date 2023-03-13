package controllerv1userregister

import (
	"errors"
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	controllers "go-clean-api/cmd/presentation/http/controllers"

	"go-clean-api/cmd/domain/entities/user"
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

		result, err := testService.Handle(controllers.HttpRequest{
			Body: dto,
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &controllers.HttpResponse{
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

		err_http := testService.HandleError(domainexceptions.InvalidEntity())
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &controllers.HttpResponseError{
			Data: controllers.HttpError{
				Code:    "INVALID_DATA",
				Message: domainexceptions.InvalidEntity().Error(),
			},
			Status: 400,
		}, err_http)
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
		err_http := testService.HandleError(domainexceptions.ErrUserAlreadyExists)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &controllers.HttpResponseError{
			Data: controllers.HttpError{
				Code:    "USER_ALREADY_EXISTS",
				Message: domainexceptions.ErrUserAlreadyExists.Error(),
			},
			Status: 400,
		}, err_http)
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
		err_http := testService.HandleError(errors.New("test"))
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
