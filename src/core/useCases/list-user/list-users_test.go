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

func Test_UseCase_ListUsers(t *testing.T) {
	t.Run("user found integration", func(t *testing.T) {
		userIntMock := newMockUserIntegration()
		mockInt := new(MockIntegration)

		mockInt.On("GetUsers").Return(userIntMock, nil)
		mockInt.AssertCalled(t, "GetUsers")

		testService := NewUseCase(mockInt)

		testService.ListUsers()
	})
}
