package api

import "strat-roulette-backend/strats"

// Start starts the REST-API on the given port with the given params
// blocks the routine it was called in
func Start(port int, stratSession strats.SessionInterface) {
	tmpSession := session{
		Strats: stratSession,
	}

	app := tmpSession.init()

	app.Listen(port)
}
