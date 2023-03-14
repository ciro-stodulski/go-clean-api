package controllerv1userregister

import (
	"errors"
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	controllers "go-clean-api/cmd/presentation/http/controllers"
	httpexceptions "go-clean-api/cmd/presentation/http/exceptions"

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

		mockUse.On("Register", dto).Return(&user.User{}, (*domainexceptions.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})

		result, errApp, err := testService.Handle(controllers.HttpRequest{
			Body: dto,
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.Nil(t, errApp)
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

		err_http := testService.HandleError(domainexceptions.InvalidEntity(), nil)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexceptions.BadRequest(controllers.HttpError{
			Code:    domainexceptions.InvalidEntity().Code,
			Message: domainexceptions.InvalidEntity().Message,
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
		err_http := testService.HandleError(domainexceptions.UserAlreadyExists(), nil)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, httpexceptions.BadRequest(controllers.HttpError{
			Code:    domainexceptions.UserAlreadyExists().Code,
			Message: domainexceptions.UserAlreadyExists().Message,
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
