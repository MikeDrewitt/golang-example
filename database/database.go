package database

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var logQuery = os.Getenv("LOG_QUERY") != ""
var connectionString = os.Getenv("CONNECTION_STRING")

func Connect() (db *gorm.DB) {
	logLevel := logger.Default.LogMode(logger.Error)

	if connectionString == "" {
		connectionString = "postgres://go_user:password@localhost:5432/go_api?sslmode=disable"
	}

	if logQuery {
		logLevel = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // makes model structs map to pluralized name User struct maps to Users (with the s) table
		},
		Logger: logLevel,
	})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("failed to initialize db connection")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
