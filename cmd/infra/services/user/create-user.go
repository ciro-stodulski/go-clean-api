package userservice

import (
	"go-api/cmd/core/entities/user"
)

func (cuuc *userService) Register(user *user.User) (*user.User, error) {

	err := cuuc.RepositoryUser.Create(user)

	return user, err
}
