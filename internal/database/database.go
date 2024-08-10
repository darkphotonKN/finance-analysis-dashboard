package database

import (
	"fmt"
	"log"
	"os"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Initializes Database and Provides DB Instance to Application

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbsslmode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbuser, dbpassword, dbhost, dbport, dbname, dbsslmode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // debug logging for sql queries
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migration
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection established and migrations run successfully.")
	return DB
}
