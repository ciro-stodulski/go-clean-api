package user

import (
	"go-clean-api/cmd/domain/entity"
	"go-clean-api/cmd/domain/exception"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID `json:"id" gorm:"unique, not null"`
	Name      string    `json:"name" gorm:"unique, not null"`
	Email     string    `json:"email" gorm:"unique, not null"`
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
		return nil, exception.InvalidEntity()
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
		return exception.InvalidEntity()
	}

	return nil
}

func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))

	if err != nil {
		return exception.PasswordIncorrect()
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
