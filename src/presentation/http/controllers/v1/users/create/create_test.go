package v1_user_create

// import (
// 	"go-api/src/core/entities/user"
// 	"go-api/src/main/container"

// 	create_dto "go-api/src/presentation/amqp/consumers/users/create/dto"
// 	ports_amqp "go-api/src/presentation/amqp/ports"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockUserCase struct {
// 	mock.Mock
// }

// func (mock *MockUserCase) CreateUser(dto create_dto.CreateDto) (*user.User, error) {
// 	arg := mock.Called(dto)
// 	result := arg.Get(0)
// 	return result.(*user.User), arg.Error(1)
// }

// func newMockUser() *user.User {
// 	user, _ := user.NewUser("test", "test", "test")
// 	return user
// }
// func Test_Consumer_User_Create(t *testing.T) {
// 	t.Run("succeffully", func(t *testing.T) {
// 		userMock := newMockUser()
// 		mockRepo := new(MockUserCase)

// 		dto := create_dto.CreateDto{
// 			Name:     "test",
// 			Email:    "test",
// 			Password: "test",
// 		}

// 		mockRepo.On("CreateUser", dto).Return(userMock, nil)

// 		testService := NewConsumer(&container.Container{
// 			CreateUserUseCase: mockRepo,
// 		})

// 		err := testService.MessageHandler(ports_amqp.Message{
// 			Body: dto,
// 		})

// 		assert.Nil(t, err)
// 		mockRepo.AssertCalled(t, "CreateUser", dto)
// 	})

// 	t.Run("return error in create use case", func(t *testing.T) {
// 		mockRepo := new(MockUserCase)

// 		dto := create_dto.CreateDto{
// 			Name:     "test",
// 			Email:    "test",
// 			Password: "test",
// 		}

// 		mockRepo.On("CreateUser", dto).Return(&user.User{}, user.ErrUserAlreadyExists)

// 		testService := NewConsumer(&container.Container{
// 			CreateUserUseCase: mockRepo,
// 		})

// 		err := testService.MessageHandler(ports_amqp.Message{
// 			Body: dto,
// 		})

// 		assert.NotNil(t, err)
// 		mockRepo.AssertCalled(t, "CreateUser", dto)
// 	})
// }
