package createuserusecasemock

import (
	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) DeleteUser(id string) error {
	arg := mock.Called(id)

	return arg.Error(1)
}
