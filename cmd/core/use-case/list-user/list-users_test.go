package listusersusecase

import (
	"go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/services/user"
	"testing"
)

func Test_UseCase_ListUsers(t *testing.T) {
	t.Run("user found integration without cache", func(t *testing.T) {
		userIntMock := mocks.NewMockUserIntegration()

		mockServices := new(mockservicesuser.MockServices)

		usecase := New(mockServices)
		mockServices.On("ListUsers").Return(userIntMock)

		usecase.ListUsers()
		mockServices.AssertCalled(t, "ListUsers")
	})
}
