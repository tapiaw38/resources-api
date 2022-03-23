package main

import (
	"github.com/tapiaw38/resources-api/database"
	"github.com/tapiaw38/resources-api/handlers"

	"log"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Connect to the database
	db := database.NewConnection()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful")

	// Initialize the handlers
	handlers.HandlerServer()
}
