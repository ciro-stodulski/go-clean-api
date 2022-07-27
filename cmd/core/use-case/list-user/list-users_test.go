package listusersusecase

import (
	"go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/infra/services/user"
	"testing"
)

func Test_UseCase_ListUsers(t *testing.T) {
	t.Run("user found integration without cache", func(t *testing.T) {
		// make mock services
		userIntMock := mocks.NewMockUserIntegration()
		mockUserServices := new(mockservicesuser.MockUserServices)
		//

		// test func
		usecase := New(mockUserServices)
		mockUserServices.On("ListUsers").Return(userIntMock)
		//

		// asserts
		usecase.ListUsers()
		mockUserServices.AssertCalled(t, "ListUsers")
		//
	})
}
