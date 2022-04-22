package main

import (
	"github.com/tapiaw38/resources-api/handlers"
	"github.com/tapiaw38/resources-api/storage"

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
	db := storage.NewConnection()

	if err := db.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful")

	// Initialize the handlers
	handlers.HandlerServer()
}
