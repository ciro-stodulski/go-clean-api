package user

import (
	entity "go-api/cmd/core/entities"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func New(email string, password string, name string) (*User, error) {
	new_u := &User{
		ID:        entity.NewID(),
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
	}

	err := new_u.Validate()

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	pwd, err := generatePassword(password)

	if err != nil {
		return nil, err
	}

	new_u.Password = pwd

	return new_u, nil
}

func (u *User) Validate() error {
	if u.Email == "" || u.Name == "" || u.Password == "" {
		return entity.ErrInvalidEntity
	}

	return nil
}

func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))

	if err != nil {
		return ErrIncorrectPassword
	}

	return nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
