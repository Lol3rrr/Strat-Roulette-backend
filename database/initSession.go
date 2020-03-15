package database

import "github.com/sirupsen/logrus"

// InitSession is used to initialize a simple Session with the given Params and connects to the Database
func InitSession(pURL, pPort, pDatabase, pCollection string) SessionInterface {
	tmpSession := &session{
		URL:        pURL,
		Port:       pPort,
		Database:   pDatabase,
		Collection: pCollection,
	}

	err := tmpSession.connect()
	if err != nil {
		logrus.Errorf("Connecting to Database: %v \n", err)
	}

	return tmpSession
}
