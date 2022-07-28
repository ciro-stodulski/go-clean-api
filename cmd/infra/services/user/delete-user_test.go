package userservice

import (
	entity "go-api/cmd/core/entities"
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
		id_mock := entity.ConvertId(userMock.ID.String())

		mockRepo.On("GetById", id_mock).Return(userMock, nil)
		mockRepo.On("DeleteById", id_mock).Return(nil)

		testService := New(mockRepo, mockIntegration, mockCache)
		err := testService.DeleteUser(userMock.ID.String())

		assert.Nil(t, err)
		mockRepo.AssertCalled(t, "GetById", id_mock)
		mockRepo.AssertCalled(t, "DeleteById", id_mock)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := mocks.NewMockUser()
		mockRepo := new(mocksqluser.MockRepository)
		mockIntegration := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)
		id_mock := entity.ConvertId(userMock.ID.String())

		errMock := user.ErrUserNotFound

		mockRepo.On("GetById", id_mock).Return(&user.User{ID: uuid.Nil}, nil)
		mockRepo.On("DeleteById", id_mock).Return(errMock)

		testService := New(mockRepo, mockIntegration, mockCache)

		err := testService.DeleteUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrUserNotFound)
		mockRepo.AssertCalled(t, "GetById", id_mock)
	})
}
