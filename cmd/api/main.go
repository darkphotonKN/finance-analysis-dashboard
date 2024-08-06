package main

import (
	// "fmt"
	"log"
	// "os"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load Environmental Variables
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := routes.SetupRouter()

	// run the server
	// port := os.Getenv("APP_PORT")

	err = router.Run(":8000")

	if err != nil {
		log.Fatalf("Server could not start %v", err)
	}
}
