package user

import (
	entity "go-api/src/core/entities"

	"github.com/google/uuid"
)

//Service  interface
type Service struct {
	RepositoryUser Repository
}

//NewService create new use case
func NewService(repository Repository) *Service {
	return &Service{
		RepositoryUser: repository,
	}
}

//GetUser Get an user
func (service *Service) GetUser(id entity.ID) (*entity.User, error) {
	// user := &entity.User{
	// 	ID:        entity.NewID(),
	// 	Name:      "oloco",
	// 	Email:     "oloco@oloco.com",
	// 	Password:  "213216574894",
	// 	CreatedAt: time.Now(),
	// }
	user, err := service.RepositoryUser.GetById(id)

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, entity.ErrUserNotFound
	}

	return user, nil
}
