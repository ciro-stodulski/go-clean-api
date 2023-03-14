package getuserusecasemock

import (
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) GetUser(id string) (*user.User, *domainexceptions.ApplicationException, error) {
	arg := mock.Called(id)
	result := arg.Get(0)

	return result.(*user.User), arg.Get(1).(*domainexceptions.ApplicationException), arg.Error(2)
}
