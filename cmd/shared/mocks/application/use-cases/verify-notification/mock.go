package verifynotificationusecasemock

import (
	portsservice "go-clean-api/cmd/domain/services"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) Notify(dto portsservice.Dto) error {
	arg := mock.Called(dto)

	return arg.Error(0)
}
