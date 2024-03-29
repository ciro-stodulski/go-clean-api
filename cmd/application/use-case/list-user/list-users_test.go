package listusersusecase

import (
	"go-clean-api/cmd/shared/mocks"
	mockservicesuser "go-clean-api/cmd/shared/mocks/application/service/user"
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
		mockUserServices.On("ListUsers").Return(userIntMock, nil)
		//

		// asserts
		usecase.Perform(nil)
		mockUserServices.AssertCalled(t, "ListUsers")
		//
	})
}
