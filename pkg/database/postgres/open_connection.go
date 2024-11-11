package postgres

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
)

func OpenConnection() (*sql.DB, error) {
	host := os.Getenv(DB_HOST)
	port := os.Getenv(DB_PORT)
	user := os.Getenv(DB_USER)
	password := os.Getenv(DB_PASSWORD)

	pgConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", pgConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
