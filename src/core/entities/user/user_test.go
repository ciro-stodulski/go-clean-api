package user_test

import (
	entity "go-api/src/core/entities/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test for create new user
func TestUserEntity(t *testing.T) {
	t.Run("create new user with succeffully", func(t *testing.T) {
		fake_name := "Test"
		fake_email := "test@test.com"
		fake_password := "1234"

		new_user_fake, err := entity.NewUser(fake_email, fake_password, fake_name)

		assert.Nil(t, err)
		assert.NotNil(t, new_user_fake.ID)
		assert.Equal(t, new_user_fake.Name, fake_name)
		assert.Equal(t, new_user_fake.Email, fake_email)
		assert.NotEqual(t, new_user_fake.Password, fake_password)
	})
	t.Run("Test for validate password with successfully new use", func(t *testing.T) {
		fake_name := "Test"
		fake_email := "test@test.com"
		fake_password := "1234"

		new_user_fake, _ := entity.NewUser(fake_email, fake_password, fake_name)

		err := new_user_fake.ValidatePassword(fake_password)

		assert.Nil(t, err)
	})
	t.Run("Test for validate password with error new user", func(t *testing.T) {
		fake_name := "Test"
		fake_email := "test@test.com"
		fake_password := "1234"

		new_user_fake, _ := entity.NewUser(fake_email, fake_password, fake_name)

		err := new_user_fake.ValidatePassword("wrong_password")

		assert.NotNil(t, err)
		assert.Equal(t, err, entity.ErrIncorrectPassword)
	})
}
