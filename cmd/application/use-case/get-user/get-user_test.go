package getuserusecase

import (
	"go-clean-api/cmd/domain/exception"
	mocks "go-clean-api/cmd/shared/mocks"
	mockservicesuser "go-clean-api/cmd/shared/mocks/infra/service/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock services
		mockServices := new(mockservicesuser.MockUserServices)
		userMock := mocks.NewMockUser()

		mockServices.On("GetUser", userMock.ID.String()).Return(userMock, (*exception.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(mockServices)
		result, err := testService.Perform(userMock.ID.String())
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
