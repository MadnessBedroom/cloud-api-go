package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQL_DB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	SQL_DB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbName string) {
	dbName = DB.Migrator().CurrentDatabase()
	return
}
