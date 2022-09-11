package usersjsonplaceholdercache

import (
	"encoding/json"
	response_jsonplaceholder "go-clean-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	mockclientcache "go-clean-api/cmd/shared/mocks/infra/repositories/cache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newMockUsers() string {
	return `[{"id": 1,"name": "Leanne Graham","username": "Bret","email": "Sincere@april.biz","address": {"street": "Kulas Light","suite": "Apt. 556","city": "Gwenborough","zipcode": "92998-3874","geo": { "lat": "-37.3159","lng":"81.1496"}},"phone": "1-770-736-8031 x56442","website": "hildegard.org","company": {"name": "Romaguera-Crona","catchPhrase": "Multi-layered client-server neural-net","bs": "harness real-time e-markets"}}]`
}

func Test_UsersCache_Get(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		userMock := newMockUsers()
		mockCache := new(mockclientcache.MockCache)

		var users []response_jsonplaceholder.User
		_ = json.Unmarshal([]byte(userMock), &users)

		mockCache.On("Get", "users").Return(userMock, nil)
		//

		// test func
		testService := New(mockCache)
		result, _ := testService.Get("users")
		//

		// asserts
		assert.NotNil(t, result)
		assert.Equal(t, result, users)
		mockCache.AssertCalled(t, "Get", "users")
		//
	})
}

func Test_UsersCache_Set(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUsers()
		mockCache := new(mockclientcache.MockCache)

		var users []response_jsonplaceholder.User
		_ = json.Unmarshal([]byte(userMock), &users)

		out, _ := json.Marshal(users)

		mockCache.On("Set", "users", string(out), 100).Return(nil)
		testService := New(mockCache)

		testService.Set("users", users, 100)

		mockCache.AssertCalled(t, "Set", "users", string(out), 100)
	})
}
