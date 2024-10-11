package driver

import (
	"yoomall/config"

	"github.com/charmbracelet/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewEmptyDB() *DB {
	return &DB{}
}

func NewDB(dsn string) *DB {
	log.Info("dsn: " + dsn)
	loggerMode := logger.Silent
	if config.Config.DEBUG {
		loggerMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(loggerMode),
	})
	if err != nil {
		panic(err)
	}
	return &DB{
		DB: db,
	}
}

func NewPostgresDB(dsn string) *DB {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{
		DB: db,
	}
}
