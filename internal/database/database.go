package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// Initializes Database and Provides DB Instance to Application

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}
