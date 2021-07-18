package entity

import "errors"

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//Password incorrect
var ErrIncorrectPassword = errors.New("password incorrect")

//Not found user error
var ErrUserNotFound = errors.New("user not found")
