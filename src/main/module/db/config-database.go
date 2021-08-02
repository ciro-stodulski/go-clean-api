package database

import (
	repositoryUser "go-api/src/infra/repositories/user"

	"github.com/jinzhu/gorm"
)

func LoadMigrationByRepositores(db *gorm.DB) {
	repositoryUser.InitMigrate(db)
}
