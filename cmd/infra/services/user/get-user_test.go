package userservice

import (
	"errors"
	"go-api/cmd/core/entities/user"
	mocks "go-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-api/cmd/shared/mocks/integrations/http/jsonplaceholder"
	mockusercache "go-api/cmd/shared/mocks/repositories/cache/user"
	mocksqluser "go-api/cmd/shared/mocks/repositories/sql/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Service_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		mockRepo.On("GetById").Return(userMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		result, err := testService.GetUser(userMock.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := mocks.NewMockUser()

		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		errMock := errors.New("err")

		mockRepo.On("GetById").Return(userMock, errMock)

		testService := New(mockRepo, mockInt, mockCache)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, errMock)
	})

	t.Run("user found integration", func(t *testing.T) {
		userIntMock := mocks.NewMockUserIntegration()

		userMockResult := &user.User{ID: uuid.Nil}
		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		mockRepo.On("GetById").Return(userMockResult, nil)
		mockInt.On("GetUsers").Return(userIntMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		result, _ := testService.GetUser("12")

		assert.NotNil(t, result)
		assert.Equal(t, userIntMock[0].Name, result.Name)
		assert.Equal(t, userIntMock[0].Email, result.Email)
		assert.Equal(t, "test_for_integration", result.Password)
	})

	t.Run("error user not found", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		userIntMock := mocks.NewMockUserIntegration()

		userMockResult := &user.User{ID: uuid.Nil}
		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		mockRepo.On("GetById").Return(userMockResult, nil)
		mockInt.On("GetUsers").Return(userIntMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrUserNotFound)
	})
}
