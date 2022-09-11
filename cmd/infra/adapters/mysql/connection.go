package msyqladapter

import (
	"fmt"
	"go-clean-api/cmd/shared/env"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(
		env.Env().DBDrive, mountConnectionString()+"&parseTime=true",
	)

	if err == nil {
		err = database.DB().Ping()
	}

	return database, err
}

func mountConnectionString() string {
	port, _ := strconv.Atoi(env.Env().DBPort)

	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8",
		env.Env().DBUser,
		env.Env().DBPassword,
		env.Env().DBHost,
		port,
		env.Env().DBSchema,
	)
}
