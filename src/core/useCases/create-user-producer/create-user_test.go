package create_user_producer

import (
	user "go-api/src/core/entities/user"
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (mock *MockProducer) CreateUser(dto create_dto.CreateDto) error {
	arg := mock.Called()
	return arg.Error(1)
}

func Test_UseCase_CreateUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockRepo := new(MockProducer)
		userMockResult := &user.User{ID: uuid.Nil}

		mockRepo.On("GetByEmail").Return(userMockResult, nil)
		mockRepo.On("Create")

		testService := NewUseCase(mockRepo)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		err := testService.CreateUser(dto)

		assert.Nil(t, err)
	})
}
