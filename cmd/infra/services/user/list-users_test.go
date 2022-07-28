package userservice

import (
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	mocks "go-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-api/cmd/shared/mocks/infra/integrations/http/jsonplaceholder"
	mockusercache "go-api/cmd/shared/mocks/infra/repositories/cache/user"
	mocksqluser "go-api/cmd/shared/mocks/infra/repositories/sql/user"
	"testing"
)

func Test_Service_ListUsers(t *testing.T) {
	t.Run("user found integration without cache", func(t *testing.T) {
		userIntMock := mocks.NewMockUserIntegration()
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockRepo := new(mocksqluser.MockRepository)
		mockCache := new(mockusercache.MockCache)

		mockCache.On("Get", "users").Return([]response_jsonplaceholder.User{}, nil)
		mockInt.On("GetUsers", 0).Return(userIntMock, nil)
		mockCache.On("Set", "users", userIntMock, 100)

		testService := New(mockRepo, mockInt, mockCache)

		testService.ListUsers()
		mockInt.AssertCalled(t, "GetUsers", 0)
		mockCache.AssertCalled(t, "Get", "users")
		mockCache.AssertCalled(t, "Set", "users", userIntMock, 100)
	})

	t.Run("user found in cache", func(t *testing.T) {
		userIntMock := mocks.NewMockUserIntegration()
		mockInt := new(mockhttpjsonplaceholder.MockIntegration)
		mockRepo := new(mocksqluser.MockRepository)
		mockCache := new(mockusercache.MockCache)

		mockCache.On("Get", "users").Return(userIntMock, nil)

		testService := New(mockRepo, mockInt, mockCache)

		testService.ListUsers()
		mockCache.AssertCalled(t, "Get", "users")
		mockCache.AssertNotCalled(t, "Get")
		mockCache.AssertNotCalled(t, "Set")
	})
}
