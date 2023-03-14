package userservice

import (
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
)

func (cuuc *userService) GetByEmail(email string) (*user.User, *domainexceptions.ApplicationException, error) {
	u, err := cuuc.SqlUser.GetByEmail(email)

	return u, nil, err
}
