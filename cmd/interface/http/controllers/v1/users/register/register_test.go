package controllerv1userregister

import (
	"errors"
	"go-api/cmd/core/entities/user"
	registeruserusecase "go-api/cmd/core/use-case/register-user"
	ports_http "go-api/cmd/interface/http/ports"
	"go-api/cmd/main/container"
	createuserusecasemock "go-api/cmd/shared/mocks/core/use-cases/create-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_User_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockUse := new(createuserusecasemock.MockUserCase)

		dto := registeruserusecase.Dto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockUse.On("Register", dto).Return(&user.User{}, nil)

		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})

		result, err := testService.Handle(ports_http.HttpRequest{
			Body: dto,
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &ports_http.HttpResponse{
			Status: 201,
		}, result)
	})

	t.Run("internal error", func(t *testing.T) {
		mockUse := new(createuserusecasemock.MockUserCase)

		dto := registeruserusecase.Dto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockUse.On("Register", dto).Return(&user.User{}, errors.New(""))

		testService := New(&container.Container{
			RegisterUserUseCase: mockUse,
		})

		result, err := testService.Handle(ports_http.HttpRequest{
			Body: dto,
		})

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}, err)
	})
}
