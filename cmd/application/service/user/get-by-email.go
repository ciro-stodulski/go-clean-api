package userservice

import (
	"go-clean-api/cmd/domain/entity/user"
)

func (cuuc *userService) GetByEmail(email string) (*user.User, error) {
	u, err := cuuc.SqlUser.GetByEmail(email)

	return u, err
}
