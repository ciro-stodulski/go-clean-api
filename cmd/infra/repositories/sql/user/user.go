package usersql

import (
	"errors"
	entity_root "go-api/cmd/core/entities"
	entity "go-api/cmd/core/entities/user"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mysqlErr *mysql.MySQLError

type (
	UserSql interface {
		DeleteById(id entity_root.ID) error
		GetById(id entity_root.ID) (*entity.User, error)
		GetByEmail(email string) (*entity.User, error)
		Create(u *entity.User) error
	}

	userSql struct {
		db *gorm.DB
	}
)

func InitMigrate(db *gorm.DB) {
	log.Default().Println("Run migration for user")

	db.AutoMigrate(&entity.User{})
}

func NewUserRepository(db *gorm.DB) (repository UserSql) {
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
			return entity.ErrUserAlreadyExists
		}
		return err.Error
	}

	return nil
}

func (ru *userSql) DeleteById(id entity_root.ID) (er error) {
	ru.db.Where("id = ?", id).Delete(&entity.User{})
	return
}
