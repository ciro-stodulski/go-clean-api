package v1_user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"go-clean-api/cmd/main/container"
	ports_http "go-clean-api/cmd/presetation/http/ports"
	"go-clean-api/cmd/shared/mocks"
	getuserusecasemock "go-clean-api/cmd/shared/mocks/application/use-cases/get-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_User_Find_By_Id(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		userMock := mocks.NewMockUser()
		mockUse := new(getuserusecasemock.MockUseCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUse.On("GetUser", id).Return(userMock, nil)
		//

		// test func
		testService := New(&container.Container{
			GetUserUseCase: mockUse,
		})

		result, err := testService.Handle(ports_http.HttpRequest{
			Params: ports_http.Params{
				ports_http.Param{Key: "id", Value: id},
			},
		})
		//

		// asserts
		assert.Nil(t, err)
		assert.NotNil(t, result)
		mockUse.AssertCalled(t, "GetUser", id)
		assert.Equal(t, &ports_http.HttpResponse{
			Data:   userMock,
			Status: 200,
		}, result)
		//
	})

	t.Run("error USER_NOT_FOUND", func(t *testing.T) {
		// make mock
		mockUse := new(getuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(&container.Container{
			GetUserUseCase: mockUse,
		})
		err_http := testService.HandleError(domainexceptions.ErrUserNotFound)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "USER_NOT_FOUND",
				Message: domainexceptions.ErrUserNotFound.Error(),
			},
			Status: 404,
		}, err_http)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockUse := new(getuserusecasemock.MockUseCase)
		//

		// test func
		testService := New(&container.Container{
			GetUserUseCase: mockUse,
		})

		err_http := testService.HandleError(domainexceptions.ErrUserAlreadyExists)
		//

		// asserts
		assert.NotNil(t, err_http)
		assert.Equal(t, &ports_http.HttpResponseError{
			Data: ports_http.HttpError{
				Code:    "INTERNAL_ERROR",
				Message: "internal error",
			},
			Status: 500,
		}, err_http)
		//
	})
}
