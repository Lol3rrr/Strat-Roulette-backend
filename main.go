package main

import (
	"strat-roulette-backend/api"
	"strat-roulette-backend/database"
	"strat-roulette-backend/strats"
	"strat-roulette-backend/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting... \n")

	stratDbURL := utils.GetEnvString("stratDB_URL", "localhost")
	stratDbPort := utils.GetEnvString("stratDB_PORT", "27017")
	stratDbDatabase := utils.GetEnvString("stratDB_DATABASE", "strats")
	stratDbCollection := utils.GetEnvString("stratDB_COLLECTION", "entrys")
	port := utils.GetEnvInt("PORT", 80)

	logrus.Info("Connecting to Strat-Database... \n")
	dbSession := database.InitSession(stratDbURL, stratDbPort, stratDbDatabase, stratDbCollection)
	logrus.Info("Connect to Strat-Database \n")

	logrus.Info("Initializing Strats Session... \n")
	stratSession := strats.InitSession(dbSession)
	logrus.Info("Initialized Strats Session \n")

	go api.Start(port, stratSession)

	logrus.Info("Started \n")

	<-make(chan bool)
}
