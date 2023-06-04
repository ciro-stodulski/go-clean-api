package domainservice

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"

	"go-clean-api/cmd/domain/entities/user"
)

type (
	UserService interface {
		GetUser(id string) (*user.User, error)
		GetByEmail(email string) (*user.User, error)
		Register(*user.User) (*user.User, error)
		DeleteUser(id string) error
		ListUsers() ([]response_jsonplaceholder.User, error)
	}
)
