package main

import (
	"errors"
	"strat-roulette-backend/api"
	"strat-roulette-backend/auth"
	"strat-roulette-backend/database"
	"strat-roulette-backend/strats"
	"strat-roulette-backend/utils"
	"time"

	"github.com/Lol3rrr/cvault"
	"github.com/sirupsen/logrus"
)

func loadAdminCreds(vSession cvault.Session) (map[string]string, error) {
	resp, err := vSession.ReadData("secret/data/strat-roulette/creds")
	if err != nil {
		return nil, err
	}

	if resp.Data["data"] == nil {
		return nil, errors.New("Secret was not set or is malformed")
	}

	data, worked := resp.Data["data"].(map[string]interface{})
	if !worked {
		return nil, errors.New("Data stored in the Secret is malformed, doesnt match 'map[string]interface{}'")
	}

	name, found := data["adminUsername"].(string)
	if !found {
		return nil, errors.New("Secret does not contain 'adminUsername'")
	}

	password, found := data["adminPassword"].(string)
	if !found {
		return nil, errors.New("Secret does not contain 'adminPassword'")
	}

	return map[string]string{
		"password": password,
		"name":     name,
	}, nil
}

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

	sessionDuration := utils.GetEnvInt("sessionDuration", 1440)

	port := utils.GetEnvInt("PORT", 80)

	vSession, err := cvault.CreateSessionEnv()
	if err != nil {
		logrus.Errorf("Could not connect to vault: %s \n", err)
		return
	}

	creds, err := loadAdminCreds(vSession)
	if err != nil {
		logrus.Errorf("Could not load the Admin Credentials: %s \n", err)
		return
	}

	adminUsername := creds["name"]
	adminPassword := creds["password"]

	logrus.Info("Connecting to Strat-Database... \n")
	stratDbSession, err := database.InitSession(stratDbURL, stratDbPort, stratDbDatabase, stratDbCollection, vSession)
	if err != nil {
		logrus.Errorf("Could not connect to the Database: %s \n", err)
		return
	}
	logrus.Info("Connected to Strat-Database \n")

	logrus.Info("Initializing Strats Session... \n")
	stratSession := strats.InitSession(stratDbSession)
	logrus.Info("Initialized Strats Session \n")

	logrus.Info("Connecting to Session-Database... \n")
	sessionDbSession, err := database.InitSession(sessionDbURL, sessionDbPort, sessionDbDatabase, sessionDbCollection, vSession)
	if err != nil {
		logrus.Errorf("Could not connect to the Database: %s \n", err)
		return
	}
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
