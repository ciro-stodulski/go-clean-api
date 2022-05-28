package jsonplaceholder

import (
	"encoding/json"
	response_jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder/responses"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIntegration struct {
	mock.Mock
}

func (mock *MockIntegration) Get(url string) ([]byte, error) {
	arg := mock.Called(url)
	result := arg.Get(0)
	return result.([]byte), arg.Error(1)
}

func newMockUsers() []byte {
	return []byte(`[{"id": 1,"name": "Leanne Graham","username": "Bret","email": "Sincere@april.biz","address": {"street": "Kulas Light","suite": "Apt. 556","city": "Gwenborough","zipcode": "92998-3874","geo": { "lat": "-37.3159","lng":"81.1496"}},"phone": "1-770-736-8031 x56442","website": "hildegard.org","company": {"name": "Romaguera-Crona","catchPhrase": "Multi-layered client-server neural-net","bs": "harness real-time e-markets"}}]`)
}

func Test_JsonPlaceholderIntegration_GetUsers(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUsers()
		mockInt := new(MockIntegration)

		var usersFake []response_jsonplaceholder.User
		_ = json.Unmarshal(userMock, &usersFake)

		mockInt.On("Get", os.Getenv("JSON_PLACE_OLDER_INTEGRATION_URL")+"/users").Return(userMock, nil)
		testService := New(mockInt)

		result, err := testService.GetUsers()

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, usersFake)
		mockInt.AssertCalled(t, "Get", os.Getenv("JSON_PLACE_OLDER_INTEGRATION_URL")+"/users")
	})
}
