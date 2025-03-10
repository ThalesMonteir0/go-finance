package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
)

func OpenConnection() (*gorm.DB, error) {
	host := os.Getenv(DB_HOST)
	port := os.Getenv(DB_PORT)
	user := os.Getenv(DB_USER)
	password := os.Getenv(DB_PASSWORD)
	dbName := os.Getenv(DB_NAME)

	pgConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(pgConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
