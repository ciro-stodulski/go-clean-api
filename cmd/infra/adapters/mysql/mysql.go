package msyqladapter

import (
	"go-clean-api/cmd/shared/env"
	"log"

	"github.com/jinzhu/gorm"
)

type MysqlAdapter struct {
	*gorm.DB
}

func (ma *MysqlAdapter) CloseDB() {
	ma.Close()
}

func (ma *MysqlAdapter) ConnectToDatabase() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	ma.DB = db

	log.Default().Print(env.Env().DBDrive + ": Connection db with succeffully")

	return nil
}
