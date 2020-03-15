package database

// InitSession is used to initialize a simple Session with the given Params and connects to the Database
func InitSession(pURL, pPort, pDatabase, pCollection string) SessionInterface {
	tmpSession := &session{
		URL:        pURL,
		Port:       pPort,
		Database:   pDatabase,
		Collection: pCollection,
	}

	tmpSession.connect()

	return tmpSession
}
