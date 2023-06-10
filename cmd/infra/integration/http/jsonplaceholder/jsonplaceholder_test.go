package jsonplaceholder

import (
	"encoding/json"
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/shared/env"
	mockhttpclient "go-clean-api/cmd/shared/mocks/infra/integration/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newMockResponseUsers() []byte {
	return []byte(`[{"id": 1,"name": "Leanne Graham","username": "Bret","email": "Sincere@april.biz","address": {"street": "Kulas Light","suite": "Apt. 556","city": "Gwenborough","zipcode": "92998-3874","geo": { "lat": "-37.3159","lng":"81.1496"}},"phone": "1-770-736-8031 x56442","website": "hildegard.org","company": {"name": "Romaguera-Crona","catchPhrase": "Multi-layered client-server neural-net","bs": "harness real-time e-markets"}}]`)
}

func Test_JsonPlaceholderIntegration_GetUsers(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		userMock := newMockResponseUsers()
		mockInt := new(mockhttpclient.MockHttpClient)

		var usersFake []response_jsonplaceholder.User
		_ = json.Unmarshal(userMock, &usersFake)

		mockInt.On("Get", env.Env().JsonPlaceOlderIntegrationUrl+"/users").Return(userMock, nil)
		//

		// test func
		testService := New(mockInt)
		result, err := testService.GetUsers()
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, usersFake)
		mockInt.AssertCalled(t, "Get", env.Env().JsonPlaceOlderIntegrationUrl+"/users")
		//
	})
}
