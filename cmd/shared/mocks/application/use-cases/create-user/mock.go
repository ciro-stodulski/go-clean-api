package createuserusecasemock

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) Register(dto domaindto.Dto) (*user.User, *domainexceptions.ApplicationException, error) {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(*user.User), arg.Get(1).(*domainexceptions.ApplicationException), arg.Error(2)
}
