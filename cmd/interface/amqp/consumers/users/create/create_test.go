package create

import (
	"go-api/cmd/core/entities/user"
	registeruserusecase "go-api/cmd/core/use-case/register-user"
	"go-api/cmd/main/container"
	"go-api/cmd/shared/mocks"
	createuserusecasemock "go-api/cmd/shared/mocks/use-cases/create-user"

	ports_amqp "go-api/cmd/interface/amqp/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Consumer_User_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		mockRepo := new(createuserusecasemock.MockUserCase)

		dto := registeruserusecase.Dto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockRepo.On("Register", dto).Return(userMock, nil)

		testService := NewConsumer(&container.Container{
			RegisterUserUseCase: mockRepo,
		})

		err := testService.MessageHandler(ports_amqp.Message{
			Body: dto,
		})

		assert.Nil(t, err)
		mockRepo.AssertCalled(t, "Register", dto)
	})

	t.Run("return error in create use case", func(t *testing.T) {
		mockRepo := new(createuserusecasemock.MockUserCase)

		dto := registeruserusecase.Dto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		mockRepo.On("Register", dto).Return(&user.User{}, user.ErrUserAlreadyExists)

		testService := NewConsumer(&container.Container{
			RegisterUserUseCase: mockRepo,
		})

		err := testService.MessageHandler(ports_amqp.Message{
			Body: dto,
		})

		assert.NotNil(t, err)
		mockRepo.AssertCalled(t, "Register", dto)
	})
}
