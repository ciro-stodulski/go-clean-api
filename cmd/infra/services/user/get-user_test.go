package userservice

import (
	"errors"
	entity "go-clean-api/cmd/domain/entities"
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	mocks "go-clean-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-clean-api/cmd/shared/mocks/infra/integrations/http/jsonplaceholder"
	mockusercache "go-clean-api/cmd/shared/mocks/infra/repositories/cache/user"
	mocksqluser "go-clean-api/cmd/shared/mocks/infra/repositories/sql/user"
	"strconv"
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
		id_mock := entity.ConvertId(userMock.ID.String())

		mockRepo.On("GetById", id_mock).Return(userMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		result, err := testService.GetUser(userMock.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
		mockRepo.AssertCalled(t, "GetById", id_mock)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := mocks.NewMockUser()

		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)
		id_mock := entity.ConvertId(userMock.ID.String())

		errMock := errors.New("err")

		mockRepo.On("GetById", id_mock).Return(userMock, errMock)

		testService := New(mockRepo, mockInt, mockCache)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, errMock)
		mockRepo.AssertCalled(t, "GetById", id_mock)
	})

	t.Run("user found integration", func(t *testing.T) {
		userIntMock := mocks.NewMockUserIntegration()
		userMockResult := &user.User{ID: uuid.Nil}
		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		id_mock := entity.ConvertId(strconv.Itoa(userIntMock[0].Id))

		mockRepo.On("GetById", id_mock).Return(userMockResult, nil)
		mockInt.On("GetUsers", 0).Return(userIntMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		result, _ := testService.GetUser(strconv.Itoa(userIntMock[0].Id))

		assert.NotNil(t, result)
		assert.Equal(t, userIntMock[0].Name, result.Name)
		assert.Equal(t, userIntMock[0].Email, result.Email)
		assert.Equal(t, "test_for_integration", result.Password)
		mockRepo.AssertCalled(t, "GetById", id_mock)
		mockInt.AssertNumberOfCalls(t, "GetUsers", 1)
	})

	t.Run("error user not found", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		userIntMock := mocks.NewMockUserIntegration()

		userMockResult := &user.User{ID: uuid.Nil}
		mockRepo := new(mocksqluser.MockRepository)
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)
		id_mock := entity.ConvertId(userMock.ID.String())

		mockRepo.On("GetById", id_mock).Return(userMockResult, nil)
		mockInt.On("GetUsers", 0).Return(userIntMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, domainexceptions.UserNotFound())
		mockRepo.AssertCalled(t, "GetById", id_mock)
		mockInt.AssertNumberOfCalls(t, "GetUsers", 1)
	})
}
