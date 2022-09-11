package msyqladapter

import (
	repositoryUser "go-clean-api/cmd/infra/repositories/sql/user"

	"github.com/jinzhu/gorm"
)

func LoadMigrationByRepositores(db *gorm.DB) {
	repositoryUser.InitMigrate(db)
}
