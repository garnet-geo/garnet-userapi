package main

import (
	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/server"
)

func main() {
	// Starting database connection
	db.InitConnection()
	defer db.CloseConnection()

	server.InitServer()
}
