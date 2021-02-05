package db

import (
	"errors"
	"fmt"
	"github.com/tpmanc/databases/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func Init(env map[string]string) {
	dbHost := env["DB_HOST"]
	dbUser := env["DB_USER"]
	dbPassword := env["DB_PASSWORD"]
	dbPort := env["DB_PORT"]
	dbName := env["DB_NAME"]

	dsn :=fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow", dbHost, dbUser, dbPassword, dbPort, dbName)
	var err error
	dbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errors.New("DB connection error")
	}
	dbConnection.AutoMigrate(&models.Databases{})
}

func Get() *gorm.DB {
	return dbConnection
}