package bootstrap

import (
	"competition-backend/pkg/config"
	"competition-backend/pkg/database"
	"competition-backend/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Get("database.username"),
		config.Get("database.password"),
		config.Get("database.host"),
		config.Get("database.port"),
		config.Get("database.database"),
		config.Get("database.charset"),
	)
	dbConfig = mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		SkipInitializeWithVersion: false,
	})

	// 连接数据库，并设置 GORM 的日志模式
	database.Conn(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

}
