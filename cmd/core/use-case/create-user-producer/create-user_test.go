package createuserproducerusecase

import (
	port "go-api/cmd/core/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (mock *MockProducer) CreateUser(dto port.CreateDto) error {
	arg := mock.Called(dto)
	return arg.Error(0)
}

func Test_UseCase_CreateUserProducer(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockPro := new(MockProducer)

		dto := port.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		mockPro.On("CreateUser", dto).Return(nil)

		testService := New(mockPro)

		err := testService.CreateUser(dto)

		mockPro.AssertCalled(t, "CreateUser", dto)
		assert.Nil(t, err)
	})
}
