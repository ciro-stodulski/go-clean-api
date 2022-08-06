package usersql

import (
	entity "go-api/cmd/core/entities/user"
	"go-api/cmd/shared/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func Test_UserRepository_GetById(t *testing.T) {
	t.Run("Should get user", func(t *testing.T) {
		// make mock
		user_mock := mocks.NewMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		mock.ExpectQuery(
			"SELECT * FROM `users` WHERE (id = ?) ORDER BY `users`.`id` ASC LIMIT 1").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
					AddRow(user_mock.ID, user_mock.Name, user_mock.Email, user_mock.Password, user_mock.CreatedAt))

		repositoryUser := New(gdb)
		//

		// test func
		result, err := repositoryUser.GetById(user_mock.ID)
		//

		// asserts
		require.NoError(t, err)
		require.Equal(t, result, user_mock)
		//
	})
}

func Test_UserRepository_GetByEmail(t *testing.T) {
	t.Run("Should get user by email ", func(t *testing.T) {
		// make mock
		user_mock := mocks.NewMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)

		mock.ExpectQuery(
			"SELECT * FROM `users` WHERE (email = ?) ORDER BY `users`.`id` ASC LIMIT 1").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
					AddRow(user_mock.ID, user_mock.Name, user_mock.Email, user_mock.Password, user_mock.CreatedAt))
		//

		// test func
		repositoryUser := New(gdb)
		result, err := repositoryUser.GetByEmail(user_mock.Email)
		//

		// asserts
		require.NoError(t, err)
		require.Equal(t, result, user_mock)
		//
	})
}

func Test_UserRepository_Create(t *testing.T) {
	t.Run("Should create user", func(t *testing.T) {
		// make mock
		user_mock := mocks.NewMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		mock.ExpectBegin()
		mock.ExpectQuery(
			"INSERT INTO `users` (`id`,`name`,`email`,`password`, `created_at`) VALUES ($1,$2,$3,$4,$5)").
			WithArgs(user_mock.ID.String(), user_mock.Name, user_mock.Email, user_mock.Password, user_mock.CreatedAt.String())
		mock.ExpectCommit()
		//

		// test func
		repositoryUser := New(gdb)
		repositoryUser.Create(&entity.User{
			ID:        user_mock.ID,
			Name:      user_mock.Name,
			Email:     user_mock.Email,
			Password:  user_mock.Password,
			CreatedAt: user_mock.CreatedAt,
		})
		//
	})
}

func Test_UserRepository_DeleteById(t *testing.T) {
	t.Run("Should delete user by id", func(t *testing.T) {
		// make mock
		user_mock := mocks.NewMockUser()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		gdb, _ := gorm.Open("mysql", db)
		mock.ExpectBegin()

		mock.ExpectQuery(
			"Delete from users WHERE id = " + user_mock.ID.String())
		mock.ExpectCommit()
		//

		// test func
		repositoryUser := New(gdb)
		err_result := repositoryUser.DeleteById(user_mock.ID)
		//

		// asserts
		require.NoError(t, err_result)
		//
	})
}
