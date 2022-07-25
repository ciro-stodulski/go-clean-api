package userservice

import (
	"go-api/cmd/core/entities/user"
	mocks "go-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-api/cmd/shared/mocks/infra/integrations/http/jsonplaceholder"
	mockusercache "go-api/cmd/shared/mocks/infra/repositories/cache/user"
	mocksqluser "go-api/cmd/shared/mocks/infra/repositories/sql/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Service_DeleteUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		mockRepo := new(mocksqluser.MockRepository)
		mockIntegration := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		mockRepo.On("GetById").Return(userMock, nil)
		mockRepo.On("DeleteById").Return(nil)

		testService := New(mockRepo, mockIntegration, mockCache)

		err := testService.DeleteUser(userMock.ID.String())

		assert.Nil(t, err)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		mockRepo := new(mocksqluser.MockRepository)
		mockIntegration := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		errMock := user.ErrUserNotFound

		mockRepo.On("GetById").Return(&user.User{ID: uuid.Nil}, nil)

		mockRepo.On("DeleteById").Return(errMock)

		testService := New(mockRepo, mockIntegration, mockCache)

		err := testService.DeleteUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrUserNotFound)
	})
}
