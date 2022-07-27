package createuserusecasemock

import (
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) DeleteUser(id string) error {
	arg := mock.Called(id)

	return arg.Error(1)
}
