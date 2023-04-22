package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Conn(config gorm.Dialector, _logger gormLogger.Interface) {
	var err error
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
