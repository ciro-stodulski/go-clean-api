package createuserusecasemock

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) Register(dto domaindto.Dto) (*user.User, error) {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(*user.User), arg.Error(1)
}
