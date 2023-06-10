package mocknosqlnotification

import (
	domaindto "go-clean-api/cmd/domain/dto"

	"github.com/stretchr/testify/mock"
)

type MockCollection struct {
	mock.Mock
}

func (mock *MockCollection) FindById(id string) (domaindto.Event, error) {
	arg := mock.Called(id)
	result := arg.Get(0)
	return result.(domaindto.Event), arg.Error(1)
}

func (mock *MockCollection) Create(id string) error {
	arg := mock.Called(id)
	return arg.Error(1)
}
