package delete_user

import (
	"errors"
	entity_root "go-api/src/core/entities"
	user "go-api/src/core/entities/user"
	"testing"

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

func (mock *MockRepository) DeleteById(id entity_root.ID) error {
	arg := mock.Called()
	return arg.Error(0)
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

func (mock *MockRepository) Create(user *user.User) {
	mock.Called()
}

func Test_UseCase_DeleteUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUser()
		mockRepo := new(MockRepository)

		mockRepo.On("DeleteById").Return(nil)

		testService := NewUseCase(mockRepo)

		err := testService.DeleteUser(userMock.ID.String())

		assert.Nil(t, err)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := newMockUser()

		mockRepo := new(MockRepository)

		errMock := errors.New("err")

		mockRepo.On("DeleteById").Return(errMock)

		testService := NewUseCase(mockRepo)

		err := testService.DeleteUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, errMock)
	})
}
