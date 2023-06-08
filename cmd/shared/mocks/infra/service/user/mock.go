package mockservicesuser

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"

	"github.com/stretchr/testify/mock"
)

type MockUserServices struct {
	mock.Mock
}

func (mock *MockUserServices) GetByEmail(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(0)
}

func (mock *MockUserServices) GetUser(id string) (*user.User, error) {
	arg := mock.Called(id)
	result := arg.Get(0)
	if result == nil {
		return nil, arg.Error(0)
	}

	return result.(*user.User), nil
}

func (mock *MockUserServices) ListUsers() ([]response_jsonplaceholder.User, error) {
	arg := mock.Called()
	result := arg.Get(0)

	return result.([]response_jsonplaceholder.User), arg.Error(1)
}

func (mock *MockUserServices) DeleteUser(id string) error {
	arg := mock.Called(id)

	return arg.Error(0)
}

func (mock *MockUserServices) Register(u *user.User) (*user.User, error) {
	arg := mock.Called(u)
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}
