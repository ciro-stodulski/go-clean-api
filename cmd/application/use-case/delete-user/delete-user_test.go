package deleteuserusecase

import (
	"go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/shared/mocks"
	mockservicesuser "go-clean-api/cmd/shared/mocks/infra/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_DeleteUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock services
		mockServices := new(mockservicesuser.MockUserServices)
		userMock := mocks.NewMockUser()

		mockServices.On("DeleteUser", userMock.ID.String()).Return((*exception.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(mockServices)
		err := testService.DeleteUser(userMock.ID.String())
		//

		// asserts
		assert.Nil(t, err)
		mockServices.AssertCalled(t, "DeleteUser", userMock.ID.String())
		//
	})
}
