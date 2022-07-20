package database

import (
	"go-api/src/shared/env"
	"log"

	"github.com/jinzhu/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (database *Database) CloseDB() {
	database.Db.Close()
}

func (database *Database) ConnectToDatabase() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	database.Db = db

	log.Default().Print(env.Env().DBDrive + ": Connection db with succeffully")

	LoadMigrationByRepositores(db)

	return nil
}
