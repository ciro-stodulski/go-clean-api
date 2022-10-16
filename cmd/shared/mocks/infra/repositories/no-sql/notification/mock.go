package mocknosqlnotification

import (
	portsservice "go-clean-api/cmd/core/ports"

	"github.com/stretchr/testify/mock"
)

type MockCollection struct {
	mock.Mock
}

func (mock *MockCollection) FindById(id string) (portsservice.Dto, error) {
	arg := mock.Called(id)
	result := arg.Get(0)
	return result.(portsservice.Dto), arg.Error(1)
}

func (mock *MockCollection) Create(id string) error {
	arg := mock.Called(id)
	return arg.Error(1)
}
