package registeruserusecase

import (
	mocks "go-api/cmd/shared/mocks"
	mockservicesnotification "go-api/cmd/shared/mocks/services/notification"
	mockservicesuser "go-api/cmd/shared/mocks/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_RegisterUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockUserServices := new(mockservicesuser.MockUserServices)
		mockNotificationServices := new(mockservicesnotification.MockNotificationServices)
		userMock := mocks.NewMockUser()

		mockUserServices.On("Register").Return(userMock, nil)

		usecase := New(mockUserServices, mockNotificationServices)

		dto := Dto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := usecase.Register(dto)

		assert.Nil(t, err)
		assert.Equal(t, userMock, result)
	})
}
