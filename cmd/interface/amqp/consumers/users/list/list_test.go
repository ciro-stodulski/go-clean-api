package list

import (
	"go-api/cmd/main/container"

	ports_amqp "go-api/cmd/interface/amqp/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) ListUsers() {
	mock.Called()
}

func Test_Consumer_List(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockRepo := new(MockUserCase)

		mockRepo.On("ListUsers").Return(nil)

		consumer := NewConsumer(&container.Container{
			ListUsersUseCase: mockRepo,
		})

		err := consumer.MessageHandler(ports_amqp.Message{Body: nil})

		assert.Nil(t, err)
		mockRepo.AssertCalled(t, "ListUsers")
	})
}
