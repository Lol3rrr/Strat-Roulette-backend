package auth

import "strat-roulette-backend/database"

// InitSession is used to create new Auth-Session
func InitSession(dbSession database.SessionInterface, pAdminUsername, pAdminPassword string, pSessionDuration int64) SessionInterface {
	sessionDuration := pSessionDuration * 60

	result := &session{
		Database:        dbSession,
		AdminUsername:   pAdminUsername,
		AdminPassword:   pAdminPassword,
		SessionDuration: sessionDuration,
	}

	return result
}
