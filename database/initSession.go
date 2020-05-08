package database

import (
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

	err := tmpSession.loadCredsAndReconnect()
	if err != nil {
		return nil, err
	}

	err = tmpSession.connect()
	if err != nil {
		logrus.Errorf("Connecting to Database: %v \n", err)
	}

	return tmpSession, nil
}
