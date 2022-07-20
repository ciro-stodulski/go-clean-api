package database

import (
	repositoryUser "go-api/src/infra/repositories/sql/user"

	"github.com/jinzhu/gorm"
)

func LoadMigrationByRepositores(db *gorm.DB) {
	repositoryUser.InitMigrate(db)
}
