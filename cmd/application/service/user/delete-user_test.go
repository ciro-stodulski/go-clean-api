package userservice

import (
	entity "go-clean-api/cmd/domain/entity"
	"go-clean-api/cmd/domain/entity/user"
	exception "go-clean-api/cmd/domain/exception"
	mocks "go-clean-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-clean-api/cmd/shared/mocks/infra/integration/http/jsonplaceholder"
	mockusercache "go-clean-api/cmd/shared/mocks/infra/repository/cache/user"
	mocksqluser "go-clean-api/cmd/shared/mocks/infra/repository/sql/user"
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

		errMock := exception.UserNotFound()

		mockRepo.On("GetById", id_mock).Return(&user.User{ID: uuid.Nil}, nil)
		mockRepo.On("DeleteById", id_mock).Return(errMock)

		testService := New(mockRepo, mockIntegration, mockCache)

		err := testService.DeleteUser(userMock.ID.String())

		assert.Equal(t, err, exception.UserNotFound())
		mockRepo.AssertCalled(t, "GetById", id_mock)
	})
}
