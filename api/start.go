package api

import (
	"strat-roulette-backend/auth"
	"strat-roulette-backend/strats"

	"github.com/sirupsen/logrus"
)

// Start starts the REST-API on the given port with the given params
// blocks the routine it was called in
func Start(port int, stratSession strats.SessionInterface, authSession auth.SessionInterface) {
	tmpSession := session{
		Strats: stratSession,
		Auth:   authSession,
	}

	app := tmpSession.init()

	err := app.Listen(port)
	if err != nil {
		logrus.Errorf("API-Listen: %v \n", err)
	}
}
