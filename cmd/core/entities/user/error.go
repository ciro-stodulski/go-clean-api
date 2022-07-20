package user

import "errors"

var ErrIncorrectPassword = errors.New("password incorrect")

var ErrUserNotFound = errors.New("user not found")

var ErrUserAlreadyExists = errors.New("user already exists")
