package main

import (
	"strat-roulette-backend/api"
	"strat-roulette-backend/auth"
	"strat-roulette-backend/database"
	"strat-roulette-backend/strats"
	"strat-roulette-backend/utils"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting... \n")

	stratDbURL := utils.GetEnvString("stratDB_URL", "localhost")
	stratDbPort := utils.GetEnvString("stratDB_PORT", "27017")
	stratDbDatabase := utils.GetEnvString("stratDB_DATABASE", "strats")
	stratDbCollection := utils.GetEnvString("stratDB_COLLECTION", "entrys")

	sessionDbURL := utils.GetEnvString("sessionDB_URL", "localhost")
	sessionDbPort := utils.GetEnvString("sessionDB_PORT", "27017")
	sessionDbDatabase := utils.GetEnvString("sessionDB_DATABASE", "strat_session")
	sessionDbCollection := utils.GetEnvString("sessionDB_COLLECTION", "sessions")

	adminUsername := utils.GetEnvString("adminUsername", "admin")
	adminPassword := utils.GetEnvString("adminPassword", "password")
	sessionDuration := utils.GetEnvInt("sessionDuration", 1440)

	port := utils.GetEnvInt("PORT", 80)

	logrus.Info("Connecting to Strat-Database... \n")
	stratDbSession := database.InitSession(stratDbURL, stratDbPort, stratDbDatabase, stratDbCollection)
	logrus.Info("Connected to Strat-Database \n")

	logrus.Info("Initializing Strats Session... \n")
	stratSession := strats.InitSession(stratDbSession)
	logrus.Info("Initialized Strats Session \n")

	logrus.Info("Connecting to Session-Database... \n")
	sessionDbSession := database.InitSession(sessionDbURL, sessionDbPort, sessionDbDatabase, sessionDbCollection)
	logrus.Info("Connected to Session-Databse \n")

	logrus.Info("Initializing Auth Session... \n")
	authSession := auth.InitSession(sessionDbSession, adminUsername, adminPassword, int64(sessionDuration))
	logrus.Info("Initialized Auth Session \n")

	utils.Schedule(func() {
		now := time.Now().Unix()

		authSession.CleanUpSessions(now)
	}, 1*time.Hour)

	go api.Start(port, stratSession, authSession)

	logrus.Info("Started \n")

	<-make(chan bool)
}
