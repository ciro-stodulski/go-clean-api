package repository

import (
	entity "go-api/src/core/entities/user"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func newMockUser() *entity.User {
	user, _ := entity.NewUser("test", "test", "test")
	return user
}

func Test_GetById(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		user_mock := newMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		repositoryUser := NewUserModel(gdb)

		mock.ExpectQuery(
			"SELECT * FROM `users` WHERE (id = ?) ORDER BY `users`.`id` ASC LIMIT 1").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
					AddRow(user_mock.ID, user_mock.Name, user_mock.Email, user_mock.Password, user_mock.CreatedAt))

		result, err := repositoryUser.GetById(user_mock.ID)
		require.NoError(t, err)
		require.Equal(t, result, user_mock)
	})
}
