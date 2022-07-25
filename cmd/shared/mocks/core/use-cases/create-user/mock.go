package createuserusecasemock

import (
	"go-api/cmd/core/entities/user"
	registeruserusecase "go-api/cmd/core/use-case/register-user"

	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) Register(dto registeruserusecase.Dto) (*user.User, error) {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(*user.User), arg.Error(1)
}
