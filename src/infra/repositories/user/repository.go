package repository

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
	"log"

	"github.com/jinzhu/gorm"
)

type (
	repositoryUser struct {
		db *gorm.DB
	}

	RepositoryUser interface {
		GetById(id entity_root.ID) (user *entity.User, er error)
	}
)

func NewUserRepository(db *gorm.DB) (repository RepositoryUser) {
	return &repositoryUser{db}
}

func InitMigrate(db *gorm.DB) {
	log.Default().Println("Run migration for user")

	db.AutoMigrate(&entity.User{})
}

func (repository *repositoryUser) GetById(id entity_root.ID) (user *entity.User, er error) {
	user = &entity.User{}
	repository.db.First(user, "id = ?", id)
	return
}
