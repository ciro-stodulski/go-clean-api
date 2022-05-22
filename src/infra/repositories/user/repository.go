package repository

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
	ports "go-api/src/core/ports"
	"log"

	"github.com/jinzhu/gorm"
)

type (
	repositoryUser struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) (repository ports.UserRepository) {
	return &repositoryUser{db}
}

func InitMigrate(db *gorm.DB) {
	log.Default().Println("Run migration for user")

	db.AutoMigrate(&entity.User{})
}

func (repository *repositoryUser) DeleteById(id entity_root.ID) (er error) {
	repository.db.Where("id = ?", id).Delete(&entity.User{})
	return
}

func (repository *repositoryUser) GetById(id entity_root.ID) (user *entity.User, er error) {
	user = &entity.User{}
	repository.db.First(user, "id = ?", id)
	return
}

func (repository *repositoryUser) GetByEmail(email string) (user *entity.User, er error) {
	user = &entity.User{}
	repository.db.First(user, "email = ?", email)
	return
}

func (repository *repositoryUser) Create(user *entity.User) {
	repository.db.Create(user)
}
