package main

import (
	"strat-roulette-backend/database"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting... \n")

	dbURL := ""
	dbPort := ""
	dbDatabase := ""
	dbCollection := ""

	log.Info("Connecting to Database... \n")
	dbSession := database.InitSession(dbURL, dbPort, dbDatabase, dbCollection)
	log.Info("Connect to Database \n")

	log.Info("Started \n")
}
