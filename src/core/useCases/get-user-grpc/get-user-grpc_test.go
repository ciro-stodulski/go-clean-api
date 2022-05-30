package get_user_grpc

import (
	"errors"
	user "go-api/src/core/entities/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newMockUser() *user.User {
	user, _ := user.NewUser("test", "test", "test")
	return user
}

type MockService struct {
	mock.Mock
}

func (mock *MockService) GetUser(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func Test_UseCase_GetUser_GRPC(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUser()
		mockService := new(MockService)

		mockService.On("GetUser").Return(userMock, nil)

		testService := NewUseCase(mockService)

		result, err := testService.GetUser(userMock.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := newMockUser()

		mockService := new(MockService)

		errMock := errors.New("err")

		mockService.On("GetUser").Return(userMock, errMock)

		testService := NewUseCase(mockService)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, errMock)
	})

	t.Run("error user not found", func(t *testing.T) {
		userMock := newMockUser()
		mockService := new(MockService)
		userMockResult := &user.User{ID: uuid.Nil}

		mockService.On("GetUser").Return(userMockResult, nil)

		testService := NewUseCase(mockService)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrUserNotFound)
	})
}
