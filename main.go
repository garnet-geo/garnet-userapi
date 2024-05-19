package main

import (
	"os"

	"github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Debugln("Starting server...")

	env := os.Getenv(consts.GlobalDebugEnv)
	if env == "" {
		log.SetLevel(log.InfoLevel)
		log.Info("Configured logger for release")
	} else {
		log.SetLevel(log.TraceLevel)
		log.Info("Configured logger for debug")
	}

	// Starting database connection
	db.InitConnection()
	defer db.CloseConnection()

	server.InitServer()
}
