package listuserusecasemock

import (
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) ListUsers() {
	mock.Called()
}
