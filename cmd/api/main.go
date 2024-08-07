package main

import (
	// "fmt"
	"fmt"
	"log"
	"os"

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
	portNo := os.Getenv("APP_PORT")
	port := fmt.Sprintf(":%s", portNo)

	fmt.Println("port:", port)

	err = router.Run(port)

	if err != nil {
		log.Fatalf("Server could not start %v", err)
	}
}
