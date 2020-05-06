package database

import (
	"errors"

	"github.com/Lol3rrr/cvault"
	"github.com/sirupsen/logrus"
)

// InitSession is used to initialize a simple Session with the given Params and connects to the Database
func InitSession(pURL, pPort, pDatabase, pCollection string, vSession cvault.Session) (SessionInterface, error) {
	tmpSession := &session{
		URL:          pURL,
		Port:         pPort,
		Database:     pDatabase,
		Collection:   pCollection,
		VaultSession: vSession,
	}

	data, err := vSession.ReadData("database/creds/strat-roulette")
	if err != nil {
		return nil, errors.New("Could not load Credentials from Vault")
	}

	vSession.RenewDataForever(data, data.LeaseDuration)

	creds := data.Data
	username, found := creds["username"]
	if !found {
		return nil, errors.New("Vault response did not include a username")
	}
	password, found := creds["password"]
	if !found {
		return nil, errors.New("Vault response did not include a password")
	}

	tmpSession.Username = username.(string)
	tmpSession.Password = password.(string)

	err = tmpSession.connect()
	if err != nil {
		logrus.Errorf("Connecting to Database: %v \n", err)
	}

	return tmpSession, nil
}
