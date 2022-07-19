package createuserusecase

import (
	entity "go-api/src/core/entities"
	user "go-api/src/core/entities/user"
	create_dto "go-api/src/interface/amqp/consumers/users/create/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newMockUser() *user.User {
	u, _ := user.New("test", "test", "test")
	return u
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetById(id entity.ID) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockRepository) GetByEmail(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockRepository) DeleteById(id entity.ID) error {
	arg := mock.Called()
	return arg.Error(1)
}

func (mock *MockRepository) Create(u *user.User) {
	mock.Called()
}

func Test_UseCase_CreateUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mr := new(MockRepository)
		umock := &user.User{ID: uuid.Nil}

		mr.On("GetByEmail").Return(umock, nil)
		mr.On("Create")

		usecase := New(mr)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := usecase.CreateUser(dto)

		assert.Nil(t, err)
		assert.Equal(t, dto.Name, result.Name)
		assert.Equal(t, dto.Email, result.Email)
	})

	t.Run("user already exists", func(t *testing.T) {
		mockRepo := new(MockRepository)
		userMockResult := newMockUser()

		mockRepo.On("GetByEmail").Return(userMockResult, nil)
		mockRepo.On("Create")

		usecase := New(mockRepo)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := usecase.CreateUser(dto)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), user.ErrUserAlreadyExists.Error())
	})
}
