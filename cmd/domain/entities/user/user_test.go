package user

import (
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Entity_User(t *testing.T) {
	t.Run("create new user with succeffully", func(t *testing.T) {
		fake_name := "Test"
		fake_email := "test@test.com"
		fake_password := "1234"

		fake_new_u, err := New(fake_email, fake_password, fake_name)

		assert.Nil(t, err)
		assert.NotNil(t, fake_new_u.ID)
		assert.Equal(t, fake_new_u.Name, fake_name)
		assert.Equal(t, fake_new_u.Email, fake_email)
		assert.NotEqual(t, fake_new_u.Password, fake_password)
	})

	t.Run("Test for validate password with successfully new use", func(t *testing.T) {
		fake_name := "Test"
		fake_email := "test@test.com"
		fake_password := "1234"

		fake_new_u, _ := New(fake_email, fake_password, fake_name)

		err := fake_new_u.ValidatePassword(fake_password)

		assert.Nil(t, err)
	})

	t.Run("Test for validate password with error new user", func(t *testing.T) {
		fake_name := "Test"
		fake_email := "test@test.com"
		fake_password := "1234"

		fake_new_u, _ := New(fake_email, fake_password, fake_name)

		err := fake_new_u.ValidatePassword("wrong_password")

		assert.NotNil(t, err)
		assert.Equal(t, err, domainexceptions.PasswordIncorrect())
	})
}
