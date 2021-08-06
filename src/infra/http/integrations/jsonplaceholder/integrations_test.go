package jsonplaceholder

import (
	response "go-api/src/infra/http/integrations/jsonplaceholder/responses"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIntegration struct {
	mock.Mock
}

func (mock *MockIntegration) Get(url string) ([]byte, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.([]byte), arg.Error(1)
}

func newMockUsers() []response.User {
	Company := response.Company{
		Name:        "test",
		CatchPhrase: "test",
		Bs:          "test",
	}

	Address := response.Address{
		Street:  "test",
		Suite:   "test",
		City:    "test",
		Zipcode: "test",
		Geo: response.Geo{
			Lat: "12",
			Lng: "234",
		},
	}

	user := response.User{
		Id:       1,
		Name:     "tes",
		Username: "test_test",
		Email:    "test@test",
		Address:  Address,
		Phone:    "test",
		Website:  "test",
		Company:  Company,
	}

	return []response.User{
		user,
	}
}

func Test_JsonPlaceholderIntegration_GetUsers(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUsers()
		mockInt := new(MockIntegration)

		mockInt.On("Get").Return(userMock, nil)

		testService := New(mockInt, "test")

		result, err := testService.GetUsers()

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
