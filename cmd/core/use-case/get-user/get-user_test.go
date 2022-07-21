package getuserusecase

import (
	mocks "go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockServices := new(mockservicesuser.MockServices)
		userMock := mocks.NewMockUser()

		mockServices.On("GetUser").Return(userMock, nil)

		testService := New(mockServices)

		result, err := testService.GetUser(userMock.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
	})
}
