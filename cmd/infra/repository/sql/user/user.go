package usersql

import (
	"errors"
	entity_root "go-clean-api/cmd/domain/entity"
	entity "go-clean-api/cmd/domain/entity/user"
	"go-clean-api/cmd/domain/exception"
	domainusersql "go-clean-api/cmd/domain/repository/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mysqlErr *mysql.MySQLError

type (
	userSql struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) (repository domainusersql.UserSql) {
	return &userSql{db}
}

func (ru *userSql) GetById(id entity_root.ID) (user *entity.User, er error) {
	user = &entity.User{}
	ru.db.First(user, "id = ?", id)
	return
}

func (ru *userSql) GetByEmail(email string) (user *entity.User, er error) {
	user = &entity.User{}
	ru.db.First(user, "email = ?", email)
	return
}

func (ru *userSql) Create(user *entity.User) error {
	err := ru.db.Create(user)

	if err != nil {
		if errors.As(err.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return exception.UserAlreadyExists()
		}
		return err.Error
	}

	return nil
}

func (ru *userSql) DeleteById(id entity_root.ID) (er error) {
	ru.db.Where("id = ?", id).Delete(&entity.User{})
	return
}
