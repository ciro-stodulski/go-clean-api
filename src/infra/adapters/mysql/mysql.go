package database

import (
	"go-api/src/shared/env"
	"log"

	"github.com/jinzhu/gorm"
)

type MysqlAdapter struct {
	Db *gorm.DB
}

func (ma *MysqlAdapter) CloseDB() {
	ma.Db.Close()
}

func (ma *MysqlAdapter) ConnectToDatabase() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	ma.Db = db

	log.Default().Print(env.Env().DBDrive + ": Connection db with succeffully")

	LoadMigrationByRepositores(db)

	return nil
}
