package user

import "errors"

//Password incorrect
var ErrIncorrectPassword = errors.New("password incorrect")

//Not found user error
var ErrUserNotFound = errors.New("user not found")
