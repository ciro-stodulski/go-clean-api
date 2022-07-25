package deleteuserusecase

import (
	"go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_DeleteUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockServices := new(mockservicesuser.MockUserServices)
		userMock := mocks.NewMockUser()

		mockServices.On("DeleteUser").Return(nil)

		testService := New(mockServices)

		err := testService.DeleteUser(userMock.ID.String())

		assert.Nil(t, err)
	})
}
