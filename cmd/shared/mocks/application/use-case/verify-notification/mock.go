package verifynotificationusecasemock

import (
	domaindto "go-clean-api/cmd/domain/dto"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) Notify(dto domaindto.Event) error {
	arg := mock.Called(dto)

	return arg.Error(0)
}
