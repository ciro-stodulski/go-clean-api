package mockhttpjsonplaceholder

import (
	response_jsonplaceholder "go-clean-api/cmd/infra/integrations/http/jsonplaceholder/responses"

	"github.com/stretchr/testify/mock"
)

type MockIntegration struct {
	mock.Mock
}

func (mock *MockIntegration) GetUsers() ([]response_jsonplaceholder.User, error) {
	arg := mock.Called(0)
	result := arg.Get(0)

	return result.([]response_jsonplaceholder.User), arg.Error(1)
}
