package strats

import (
	"github.com/Lol3rrr/mongovault"
)

// InitSession creates a new Strats-Session with the given params
func InitSession(dbSession mongovault.DB) SessionInterface {
	tmpSession := &session{
		Database: dbSession,
	}

	return tmpSession
}
