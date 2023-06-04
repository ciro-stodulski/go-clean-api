package getuserusecasemock

import (
	"go-clean-api/cmd/domain/entities/user"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (mock *MockUseCase) GetUser(id string) (*user.User, error) {
	arg := mock.Called(id)
	result := arg.Get(0)

	return result.(*user.User), arg.Error(1)
}
