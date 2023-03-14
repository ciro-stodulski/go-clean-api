package userservice

import (
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
)

func (cuuc *userService) Register(user *user.User) (*user.User, *domainexceptions.ApplicationException, error) {

	err := cuuc.SqlUser.Create(user)

	return user, nil, err
}
