package get_user

import (
	"errors"
	entity_root "go-api/src/core/entities"
	user "go-api/src/core/entities/user"
	response_jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder/responses"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newMockUser() *user.User {
	user, _ := user.NewUser("test", "test", "test")
	return user
}

func newMockUserIntegration() []response_jsonplaceholder.User {
	return []response_jsonplaceholder.User{{
		Id:       12,
		Name:     "test",
		Username: "test",
		Email:    "test@test",
		Phone:    "test",
		Website:  "test",
		Address: response_jsonplaceholder.Address{
			Street:  "test",
			Suite:   "test",
			City:    "test",
			Zipcode: "test",
			Geo: response_jsonplaceholder.Geo{
				Lat: "test",
				Lng: "test",
			},
		},
		Company: response_jsonplaceholder.Company{
			Name:        "test",
			CatchPhrase: "test",
			Bs:          "test",
		},
	}}
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetById(id entity_root.ID) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

type MockIntegration struct {
	mock.Mock
}

func (mock *MockIntegration) GetUsers() ([]response_jsonplaceholder.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.([]response_jsonplaceholder.User), arg.Error(1)
}

func Test_UseCase_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUser()
		mockRepo := new(MockRepository)
		mockInt := new(MockIntegration)

		mockRepo.On("GetById").Return(userMock, nil)

		testService := NewUseCase(mockRepo, mockInt)

		result, err := testService.GetUser(userMock.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := newMockUser()

		mockRepo := new(MockRepository)
		mockInt := new(MockIntegration)

		errMock := errors.New("err")

		mockRepo.On("GetById").Return(userMock, errMock)

		testService := NewUseCase(mockRepo, mockInt)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, errMock)
	})

	t.Run("user found integration", func(t *testing.T) {
		userIntMock := newMockUserIntegration()

		userMockResult := &user.User{ID: uuid.Nil}
		mockRepo := new(MockRepository)
		mockInt := new(MockIntegration)

		mockRepo.On("GetById").Return(userMockResult, nil)
		mockInt.On("GetUsers").Return(userIntMock, nil)

		testService := NewUseCase(mockRepo, mockInt)

		result, err := testService.GetUser("12")

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, userIntMock[0].Name, result.Name)
		assert.Equal(t, userIntMock[0].Email, result.Email)
		assert.Equal(t, "test_for_integration", result.Password)
	})

	t.Run("error user not found", func(t *testing.T) {
		userMock := newMockUser()
		userIntMock := newMockUserIntegration()

		userMockResult := &user.User{ID: uuid.Nil}
		mockRepo := new(MockRepository)
		mockInt := new(MockIntegration)

		mockRepo.On("GetById").Return(userMockResult, nil)
		mockInt.On("GetUsers").Return(userIntMock, nil)

		testService := NewUseCase(mockRepo, mockInt)

		_, err := testService.GetUser(userMock.ID.String())

		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrUserNotFound)
	})
}
