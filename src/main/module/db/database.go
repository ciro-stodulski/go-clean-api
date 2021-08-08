package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (database *Database) CloseDB() {
	database.Db.Close()
}

func (database *Database) ConnectToDabase() error {
	db, err := GetDatabase()

	if err != nil {
		return err
	}

	database.Db = db

	log.Default().Print("connection db with succeffully")

	LoadMigrationByRepositores(db)

	return nil
}
