package deleteeuserusecasemock

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) DeleteUser(id string) (*domainexceptions.ApplicationException, error) {
	arg := mock.Called(id)
	result := arg.Get(0)

	return result.(*domainexceptions.ApplicationException), arg.Error(1)
}
