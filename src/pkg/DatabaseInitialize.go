package models

import (
	"app-test/src/pkg/config"
	"app-test/src/pkg/datasource"
	"app-test/src/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"time"
)

func InitDB(config *config.Config) {
	var gdb *gorm.DB
	var err error
	if config.DataSource.DBType == "mysql" {
		config.DataSource.DSN = config.MySQL.DSN()
	} else if config.DataSource.DBType == "sqlite3" {
		config.DataSource.DSN = config.SQLite.DSN()
	}
	logger.Info(config.DataSource.DSN)
	gdb, err = gorm.Open(config.DataSource.DBType, config.DataSource.DSN)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	if config.DataSource.Debug {
		gdb.LogMode(true)
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	gdb.DB().SetMaxIdleConns(config.DataSource.MaxIdleConnections)
	gdb.DB().SetMaxOpenConns(config.DataSource.MaxOpenConnections)
	gdb.DB().SetConnMaxLifetime(time.Duration(config.DataSource.MaxLifetime) * time.Second)
	datasource.DB = gdb
}
