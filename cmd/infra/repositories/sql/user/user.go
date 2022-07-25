package userepository

import (
	"errors"
	entity_root "go-api/cmd/core/entities"
	entity "go-api/cmd/core/entities/user"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type (
	repositoryUser struct {
		db *gorm.DB
	}
)

var mysqlErr *mysql.MySQLError

func NewUserRepository(db *gorm.DB) (repository UserRepository) {
	return &repositoryUser{db}
}

func InitMigrate(db *gorm.DB) {
	log.Default().Println("Run migration for user")

	db.AutoMigrate(&entity.User{})
}

func (ru *repositoryUser) GetById(id entity_root.ID) (user *entity.User, er error) {
	user = &entity.User{}
	ru.db.First(user, "id = ?", id)
	return
}

func (ru *repositoryUser) GetByEmail(email string) (user *entity.User, er error) {
	user = &entity.User{}
	ru.db.First(user, "email = ?", email)
	return
}

func (ru *repositoryUser) Create(user *entity.User) error {
	err := ru.db.Create(user)

	if errors.As(err.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		return entity.ErrUserAlreadyExists
	}

	return nil
}

func (ru *repositoryUser) DeleteById(id entity_root.ID) (er error) {
	ru.db.Where("id = ?", id).Delete(&entity.User{})
	return
}
