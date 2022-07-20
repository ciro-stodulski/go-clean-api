package users_cache

import (
	"encoding/json"
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (mock *MockCache) Get(key string) (string, error) {
	arg := mock.Called(key)
	result := arg.Get(0)
	return result.(string), arg.Error(1)
}

func (mock *MockCache) Set(key string, value string, timeEx int) error {
	arg := mock.Called(key, value, timeEx)
	return arg.Error(0)
}

func newMockUsers() string {
	return `[{"id": 1,"name": "Leanne Graham","username": "Bret","email": "Sincere@april.biz","address": {"street": "Kulas Light","suite": "Apt. 556","city": "Gwenborough","zipcode": "92998-3874","geo": { "lat": "-37.3159","lng":"81.1496"}},"phone": "1-770-736-8031 x56442","website": "hildegard.org","company": {"name": "Romaguera-Crona","catchPhrase": "Multi-layered client-server neural-net","bs": "harness real-time e-markets"}}]`
}

func Test_UsersCache_Get(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUsers()
		mockCache := new(MockCache)

		var users []response_jsonplaceholder.User
		_ = json.Unmarshal([]byte(userMock), &users)

		mockCache.On("Get", "users").Return(userMock, nil)
		testService := New(mockCache)

		result, _ := testService.Get("users")

		assert.NotNil(t, result)
		assert.Equal(t, result, users)
		mockCache.AssertCalled(t, "Get", "users")
	})
}

func Test_UsersCache_Set(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUsers()
		mockCache := new(MockCache)

		var users []response_jsonplaceholder.User
		_ = json.Unmarshal([]byte(userMock), &users)

		out, _ := json.Marshal(users)

		mockCache.On("Set", "users", string(out), 100).Return(nil)
		testService := New(mockCache)

		testService.Set("users", users, 100)

		mockCache.AssertCalled(t, "Set", "users", string(out), 100)
	})
}
