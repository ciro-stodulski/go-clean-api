package create_user_producer

import (
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (mock *MockProducer) CreateUser(dto create_dto.CreateDto) error {
	arg := mock.Called(dto)
	return arg.Error(0)
}

func Test_UseCase_CreateUserProducer(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockPro := new(MockProducer)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		mockPro.On("CreateUser", dto).Return(nil)

		testService := NewUseCase(mockPro)

		err := testService.CreateUser(dto)

		mockPro.AssertCalled(t, "CreateUser", dto)
		assert.Nil(t, err)
	})
}
