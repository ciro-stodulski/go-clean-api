package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//Data for user
type User struct {
	ID        ID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

//Create a new user
func NewUser(email, password, name string) (*User, error) {
	new_user := &User{
		ID:        NewID(),
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
	}

	err := new_user.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	pwd, err := generatePassword(password)

	if err != nil {
		return nil, err
	}

	new_user.Password = pwd

	return new_user, nil
}

//Validate validate props by User
func (u *User) Validate() error {
	if u.Email == "" || u.Name == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}

//Validate password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))

	if err != nil {
		return ErrIncorrectPassword
	}

	return nil
}

//Generate passord password
func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
