package registeruserusecase

import (
	mocks "go-api/cmd/shared/mocks"
	mockservicesnotification "go-api/cmd/shared/mocks/infra/services/notification"
	mockservicesuser "go-api/cmd/shared/mocks/infra/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_RegisterUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockUserServices := new(mockservicesuser.MockUserServices)
		mockNotificationServices := new(mockservicesnotification.MockNotificationServices)
		dto := Dto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		userMock := mocks.CreateMockUser(dto.Name, dto.Email, dto.Password)

		mockUserServices.On("Register").Return(userMock, nil)

		usecase := New(mockUserServices, mockNotificationServices)

		result, err := usecase.Register(dto)

		assert.Nil(t, err)
		assert.Equal(t, userMock, result)
	})
}
