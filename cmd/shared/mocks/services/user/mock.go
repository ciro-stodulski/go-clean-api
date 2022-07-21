package mockservicesuser

import (
	"go-api/cmd/core/entities/user"
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	dto "go-api/cmd/interface/amqp/consumers/users/create/dto"

	"github.com/stretchr/testify/mock"
)

type MockServices struct {
	mock.Mock
}

func (mock *MockServices) GetUser(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockServices) ListUsers() []response_jsonplaceholder.User {
	arg := mock.Called()
	result := arg.Get(0)
	return result.([]response_jsonplaceholder.User)
}

func (mock *MockServices) DeleteUser(id string) error {
	arg := mock.Called()
	return arg.Error(0)
}

func (mock *MockServices) CreateUser(dto dto.CreateDto) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}
