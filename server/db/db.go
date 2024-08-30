package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUrl := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
