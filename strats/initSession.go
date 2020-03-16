package strats

import "strat-roulette-backend/database"

// InitSession creates a new Strats-Session with the given params
func InitSession(dbSession database.SessionInterface) SessionInterface {
	tmpSession := &session{
		Database: dbSession,
	}

	return tmpSession
}
