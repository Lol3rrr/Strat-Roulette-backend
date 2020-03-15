package database

// CreateSession is used to initialize a simple Session with the given Params
func CreateSession(pURL, pPort, pDatabase, pCollection string) SessionInterface {
	return &session{
		URL:        pURL,
		Port:       pPort,
		Database:   pDatabase,
		Collection: pCollection,
	}
}
