package msyqladapter

import (
	repositoryUser "go-clean-api/cmd/infra/repository/sql/user"

	"github.com/jinzhu/gorm"
)

func LoadMigrationByRepositores(db *gorm.DB) {
	repositoryUser.InitMigrate(db)
}
