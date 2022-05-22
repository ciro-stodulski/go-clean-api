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

func Test_UserRepository_GetById(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		user_mock := newMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		repositoryUser := NewUserRepository(gdb)

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

func Test_UserRepository_GetByEmail(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		user_mock := newMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		repositoryUser := NewUserRepository(gdb)

		mock.ExpectQuery(
			"SELECT * FROM `users` WHERE (email = ?) ORDER BY `users`.`id` ASC LIMIT 1").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
					AddRow(user_mock.ID, user_mock.Name, user_mock.Email, user_mock.Password, user_mock.CreatedAt))

		result, err := repositoryUser.GetByEmail(user_mock.Email)
		require.NoError(t, err)
		require.Equal(t, result, user_mock)
	})
}

func Test_UserRepository_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		user_mock := newMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		repositoryUser := NewUserRepository(gdb)
		mock.ExpectQuery(
			"INSERT INTO `users` (`id`,`name`,`email`,`password`, `created_at`) VALUES ('" + user_mock.ID.String() + "','" + user_mock.Name + "', '" + user_mock.Email + "', '" + user_mock.Password + "', '" + user_mock.CreatedAt.String() + "')")

		repositoryUser.Create(&entity.User{
			ID:        user_mock.ID,
			Name:      user_mock.Name,
			Email:     user_mock.Email,
			Password:  user_mock.Password,
			CreatedAt: user_mock.CreatedAt,
		})

	})
}

func Test_UserRepository_DeleteById(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		user_mock := newMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		repositoryUser := NewUserRepository(gdb)

		mock.ExpectQuery(
			"Delete FROM `users` WHERE id = " + user_mock.ID.String())

		err_result := repositoryUser.DeleteById(user_mock.ID)
		require.NoError(t, err_result)
	})
}
