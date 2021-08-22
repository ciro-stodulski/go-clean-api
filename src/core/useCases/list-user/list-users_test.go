package list_users

import (
	response_jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder/responses"
	"testing"

	"github.com/stretchr/testify/mock"
)

func newMockUserIntegration() []response_jsonplaceholder.User {
	return []response_jsonplaceholder.User{{
		Id:       12,
		Name:     "test",
		Username: "test",
		Email:    "test@test",
		Phone:    "test",
		Website:  "test",
		Address: response_jsonplaceholder.Address{
			Street:  "test",
			Suite:   "test",
			City:    "test",
			Zipcode: "test",
			Geo: response_jsonplaceholder.Geo{
				Lat: "test",
				Lng: "test",
			},
		},
		Company: response_jsonplaceholder.Company{
			Name:        "test",
			CatchPhrase: "test",
			Bs:          "test",
		},
	}}
}

type MockIntegration struct {
	mock.Mock
}

func (mock *MockIntegration) GetUsers() ([]response_jsonplaceholder.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.([]response_jsonplaceholder.User), arg.Error(1)
}

type MockCache struct {
	mock.Mock
}

func (mock *MockCache) Get(key string) ([]response_jsonplaceholder.User, error) {
	arg := mock.Called(key)
	result := arg.Get(0)
	return result.([]response_jsonplaceholder.User), arg.Error(1)
}

func (mock *MockCache) Set(key string, value []response_jsonplaceholder.User, timeEx int) {
	mock.Called(key, value, timeEx)
}

func Test_UseCase_ListUsers(t *testing.T) {
	t.Run("user found integration without cache", func(t *testing.T) {
		userIntMock := newMockUserIntegration()
		mockInt := new(MockIntegration)

		mockCache := new(MockCache)

		mockCache.On("Get", "users").Return([]response_jsonplaceholder.User{}, nil)
		mockInt.On("GetUsers").Return(userIntMock, nil)
		mockCache.On("Set", "users", userIntMock, 100)

		testService := NewUseCase(mockInt, mockCache)

		testService.ListUsers()
		mockInt.AssertCalled(t, "GetUsers")
		mockCache.AssertCalled(t, "Get", "users")
		mockCache.AssertCalled(t, "Set", "users", userIntMock, 100)
	})

	t.Run("user found in cache", func(t *testing.T) {
		userIntMock := newMockUserIntegration()
		mockInt := new(MockIntegration)

		mockCache := new(MockCache)

		mockCache.On("Get", "users").Return(userIntMock, nil)

		testService := NewUseCase(mockInt, mockCache)

		testService.ListUsers()
		mockCache.AssertCalled(t, "Get", "users")
		mockCache.AssertNotCalled(t, "Get")
		mockCache.AssertNotCalled(t, "Set")
	})
}
