package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger:                                   _logger,
		SkipDefaultTransaction:                   true, // 跳过默认事务
		DisableForeignKeyConstraintWhenMigrating: true, // 外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbName string) {
	dbName = DB.Migrator().CurrentDatabase()
	return
}
