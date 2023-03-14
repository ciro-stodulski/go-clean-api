package verifynotificationusecasemock

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) Notify(dto domaindto.Event) (*domainexceptions.ApplicationException, error) {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(*domainexceptions.ApplicationException), arg.Error(0)
}
