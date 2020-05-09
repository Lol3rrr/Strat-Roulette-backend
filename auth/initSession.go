package auth

import (
	"github.com/Lol3rrr/mongovault"
)

// InitSession is used to create new Auth-Session
func InitSession(dbSession mongovault.DB, pAdminUsername, pAdminPassword string, pSessionDuration int64) SessionInterface {
	sessionDuration := pSessionDuration * 60

	result := &session{
		Database:        dbSession,
		AdminUsername:   pAdminUsername,
		AdminPassword:   pAdminPassword,
		SessionDuration: sessionDuration,
	}

	return result
}
