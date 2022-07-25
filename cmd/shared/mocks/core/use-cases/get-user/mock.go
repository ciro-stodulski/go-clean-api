package getuserusecasemock

import (
	"go-api/cmd/core/entities/user"

	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) GetUser(id string) (*user.User, error) {
	arg := mock.Called(id)
	result := arg.Get(0)

	return result.(*user.User), arg.Error(1)
}
