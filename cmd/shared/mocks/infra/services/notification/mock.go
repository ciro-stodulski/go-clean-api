package mockservicesnotification

import (
	domaindto "go-clean-api/cmd/domain/dto"

	"github.com/stretchr/testify/mock"
)

type MockNotificationServices struct {
	mock.Mock
}

func (mock *MockNotificationServices) SendNotify(dto domaindto.Event) error {
	arg := mock.Called(dto)

	return arg.Error(0)
}

func (mock *MockNotificationServices) CheckNotify(msg string) error {
	arg := mock.Called(msg)

	return arg.Error(0)
}

func (mock *MockNotificationServices) SaveNotify(dto domaindto.Event) string {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(string)
}

func (mock *MockNotificationServices) FindById(msg string) (*domaindto.Event, error) {
	arg := mock.Called(msg)
	result := arg.Get(0)

	return result.(*domaindto.Event), arg.Error(1)
}
