package database

import (
	"fmt"
	repositoryUser "go-api/src/infra/repositories/user"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DbConfig struct {
	Driver   string
	Schema   string
	Host     string
	Port     int
	Username string
	Password string
	Pool     struct {
		Min uint8
		Max uint8
	}
}

func mountConnectionString() string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_SCHEMA"),
	)
}

func LoadMigrationByRepositores(db *gorm.DB) {
	repositoryUser.InitMigrate(db)
}

func GetDatabase(cfg *DbConfig) (*gorm.DB, error) {
	database, err := gorm.Open(cfg.Driver, mountConnectionString()+"&parseTime=true")

	if err == nil {
		err = database.DB().Ping()
	}

	return database, err
}

func NewDbConfig() (cfg *DbConfig, err error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dbConfig := &DbConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Driver:   os.Getenv("DB_DRIVE"),
	}

	return dbConfig, nil
}
