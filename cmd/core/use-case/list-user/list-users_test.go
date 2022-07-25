package listusersusecase

import (
	"go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/infra/services/user"
	"testing"
)

func Test_UseCase_ListUsers(t *testing.T) {
	t.Run("user found integration without cache", func(t *testing.T) {
		userIntMock := mocks.NewMockUserIntegration()

		mockUserServices := new(mockservicesuser.MockUserServices)

		usecase := New(mockUserServices)
		mockUserServices.On("ListUsers").Return(userIntMock)

		usecase.ListUsers()
		mockUserServices.AssertCalled(t, "ListUsers")
	})
}
