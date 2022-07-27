package mockservicesuser

import (
	"go-api/cmd/core/entities/user"
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"

	"github.com/stretchr/testify/mock"
)

type MockUserServices struct {
	mock.Mock
}

func (mock *MockUserServices) GetByEmail(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockUserServices) GetUser(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockUserServices) ListUsers() []response_jsonplaceholder.User {
	arg := mock.Called()
	result := arg.Get(0)
	return result.([]response_jsonplaceholder.User)
}

func (mock *MockUserServices) DeleteUser(id string) error {
	arg := mock.Called()
	return arg.Error(0)
}

func (mock *MockUserServices) Register(u *user.User) (*user.User, error) {
	arg := mock.Called(u)
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}
