package mocksqluser

import (
	entity_root "go-clean-api/cmd/domain/entity"
	user "go-clean-api/cmd/domain/entity/user"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetById(id entity_root.ID) (*user.User, error) {
	arg := mock.Called(id)
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockRepository) GetByEmail(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func (mock *MockRepository) Create(user *user.User) error {
	arg := mock.Called()
	return arg.Error(0)
}

func (mock *MockRepository) DeleteById(id entity_root.ID) error {
	arg := mock.Called(id)
	return arg.Error(0)
}
