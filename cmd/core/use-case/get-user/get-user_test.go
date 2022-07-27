package getuserusecase

import (
	mocks "go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/infra/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock services
		mockServices := new(mockservicesuser.MockUserServices)
		userMock := mocks.NewMockUser()

		mockServices.On("GetUser", userMock.ID.String()).Return(userMock, nil)
		//

		// test func
		testService := New(mockServices)
		result, err := testService.GetUser(userMock.ID.String())
		//

		// asserts
		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
		mockServices.AssertCalled(t, "GetUser", userMock.ID.String())
		//
	})
}
