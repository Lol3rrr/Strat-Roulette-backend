package main

import (
	"strat-roulette-backend/api"
	"strat-roulette-backend/database"
	"strat-roulette-backend/strats"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting... \n")

	dbURL := ""
	dbPort := ""
	dbDatabase := ""
	dbCollection := ""
	port := 8080

	logrus.Info("Connecting to Database... \n")
	dbSession := database.InitSession(dbURL, dbPort, dbDatabase, dbCollection)
	logrus.Info("Connect to Database \n")

	logrus.Info("Initializing Strats Session... \n")
	stratSession := strats.InitSession(dbSession)
	logrus.Info("Initialized Strats Session \n")

	go api.Start(port, stratSession)

	logrus.Info("Started \n")

	<-make(chan bool)
}
