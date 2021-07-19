package model

import (
	entity "go-api/src/core/entities"

	"github.com/jinzhu/gorm"
)

type (
	repositoryUser struct {
		db *gorm.DB
	}

	RepositoryUser interface {
		GetById(id entity.ID) (user *entity.User, er error)
	}
)

func NewUserModel(db *gorm.DB) (repository RepositoryUser) {
	//db.AutoMigrate(&entity.User{})
	return &repositoryUser{db}
}

func (repository *repositoryUser) GetById(id entity.ID) (user *entity.User, er error) {
	user = &entity.User{}
	repository.db.First(user, "id = ?", id)
	return
}
