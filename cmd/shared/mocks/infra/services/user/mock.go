package mockservicesuser

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"github.com/stretchr/testify/mock"
)

type MockUserServices struct {
	mock.Mock
}

func (mock *MockUserServices) GetByEmail(id string) (*user.User, *domainexceptions.ApplicationException, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), result.(*domainexceptions.ApplicationException), arg.Error(1)
}

func (mock *MockUserServices) GetUser(id string) (*user.User, *domainexceptions.ApplicationException, error) {
	arg := mock.Called(id)
	result := arg.Get(0)
	if result == nil {
		return nil, nil, arg.Error(1)
	}

	return result.(*user.User), arg.Get(1).(*domainexceptions.ApplicationException), arg.Error(2)
}

func (mock *MockUserServices) ListUsers() ([]response_jsonplaceholder.User, *domainexceptions.ApplicationException, error) {
	arg := mock.Called()
	result := arg.Get(0)

	return result.([]response_jsonplaceholder.User), arg.Get(1).(*domainexceptions.ApplicationException), arg.Error(2)
}

func (mock *MockUserServices) DeleteUser(id string) (*domainexceptions.ApplicationException, error) {
	arg := mock.Called(id)
	result := arg.Get(0)

	return result.(*domainexceptions.ApplicationException), arg.Error(0)
}

func (mock *MockUserServices) Register(u *user.User) (*user.User, *domainexceptions.ApplicationException, error) {
	arg := mock.Called(u)
	result := arg.Get(0)
	return result.(*user.User), arg.Get(1).(*domainexceptions.ApplicationException), arg.Error(2)
}
