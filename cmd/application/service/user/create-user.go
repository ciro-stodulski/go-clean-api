package userservice

import (
	"go-clean-api/cmd/domain/entity/user"
)

func (cuuc *userService) Register(user *user.User) (*user.User, error) {

	err := cuuc.SqlUser.Create(user)

	return user, err
}
