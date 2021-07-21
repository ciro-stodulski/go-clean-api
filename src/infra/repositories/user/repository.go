package repository

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"

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

func NewUserModel(db *gorm.DB) (repository RepositoryUser) {
	db.AutoMigrate(&entity.User{})
	return &repositoryUser{db}
}

func (repository *repositoryUser) GetById(id entity_root.ID) (user *entity.User, er error) {
	user = &entity.User{}
	repository.db.First(user, "id = ?", id)
	return
}
