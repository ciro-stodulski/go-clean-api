package create_user

import (
	entity_root "go-api/src/core/entities"
	user "go-api/src/core/entities/user"
	create_dto "go-api/src/presentation/amqp/consumers/users/create/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newMockUser() *user.User {
	user, _ := user.NewUser("test", "test", "test")
	return user
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetById(id entity_root.ID) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockRepository) GetByEmail(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockRepository) DeleteById(id entity_root.ID) error {
	arg := mock.Called()
	return arg.Error(1)
}

func (mock *MockRepository) Create(user *user.User) {
	mock.Called()
}

func Test_UseCase_CreateUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockRepo := new(MockRepository)
		userMockResult := &user.User{ID: uuid.Nil}

		mockRepo.On("GetByEmail").Return(userMockResult, nil)
		mockRepo.On("Create")

		testService := NewUseCase(mockRepo)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := testService.CreateUser(dto)

		assert.Nil(t, err)
		assert.Equal(t, dto.Name, result.Name)
		assert.Equal(t, dto.Email, result.Email)
	})

	t.Run("user already exists", func(t *testing.T) {
		mockRepo := new(MockRepository)
		userMockResult := newMockUser()

		mockRepo.On("GetByEmail").Return(userMockResult, nil)
		mockRepo.On("Create")

		testService := NewUseCase(mockRepo)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := testService.CreateUser(dto)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), user.ErrUserAlreadyExists.Error())
	})
}
