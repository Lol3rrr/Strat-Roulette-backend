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
	dbSession := database.CreateSession(dbURL, dbPort, dbDatabase, dbCollection)
	err := dbSession.Connect()
	if err != nil {
		log.Errorf("Failed to connect to Database: %v \n", err)
		return
	}
	log.Info("Connect to Database \n")

	log.Info("Started \n")
}
