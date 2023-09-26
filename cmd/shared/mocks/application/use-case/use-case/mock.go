package usecasemock

import (
	"github.com/stretchr/testify/mock"
)

type MockUseCase[I any, O any] struct {
	mock.Mock
}

func (mock *MockUseCase[I, O]) Perform(id I) (O, error) {
	arg := mock.Called(id)
	result := arg.Get(0)

	return result.(O), arg.Error(1)
}
