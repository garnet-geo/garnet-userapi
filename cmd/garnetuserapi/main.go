package main

import (
	"log"

	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Starting database connection
	db.InitConnection()
	defer db.CloseConnection()

	server.InitServer()
}
