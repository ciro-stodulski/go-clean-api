package v1_user

import (
	ports_http "go-api/cmd/interface/http/ports"
	"go-api/cmd/main/container"
	"go-api/cmd/shared/mocks"
	getuserusecasemock "go-api/cmd/shared/mocks/core/use-cases/get-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller_Find_By_Id(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		mockUse := new(getuserusecasemock.MockUserCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUse.On("GetUser").Return(userMock, nil)

		testService := New(&container.Container{
			GetUserUseCase: mockUse,
		})

		result, err := testService.Handle(ports_http.HttpRequest{
			Params: ports_http.Params{
				ports_http.Param{Key: "id", Value: id},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &ports_http.HttpResponse{
			Data:   userMock,
			Status: 200,
		}, result)
	})
}
