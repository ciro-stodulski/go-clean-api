package v1_user_create

import (
	"errors"
	"go-api/src/main/container"
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
	ports_http "go-api/src/presentation/http/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) CreateUser(dto create_dto.CreateDto) error {
	arg := mock.Called(dto)
	return arg.Error(0)
}

func Test_Consumer_User_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockUse := new(MockUserCase)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockUse.On("CreateUser", dto).Return(nil)

		testService := NewController(&container.Container{
			CreateUserProducerUseCase: mockUse,
		})

		result, err := testService.LoadRoute().Handle(ports_http.HttpRequest{
			Body: dto,
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &ports_http.HttpResponse{
			Status: 200,
		}, result)
	})

	t.Run("internal error", func(t *testing.T) {
		mockUse := new(MockUserCase)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockUse.On("CreateUser", dto).Return(errors.New(""))

		testService := NewController(&container.Container{
			CreateUserProducerUseCase: mockUse,
		})

		result, err := testService.LoadRoute().Handle(ports_http.HttpRequest{
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
