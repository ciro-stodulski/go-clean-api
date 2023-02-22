package userservice

import (
	"go-clean-api/cmd/domain/entities/user"
)

func (cuuc *userService) GetByEmail(email string) (*user.User, error) {
	u, err := cuuc.SqlUser.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return u, nil
}
