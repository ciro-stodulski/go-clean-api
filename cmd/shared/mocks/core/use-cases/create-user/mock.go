package createuserusecasemock

import (
	"go-clean-api/cmd/core/entities/user"
	registeruserusecase "go-clean-api/cmd/core/use-case/register-user"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) Register(dto registeruserusecase.Dto) (*user.User, error) {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(*user.User), arg.Error(1)
}
