package api

import "strat-roulette-backend/strats"

type session struct {
	Strats strats.SessionInterface
}

type addStratRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Site        strats.Site       `json:"site"`
	Modes       []strats.GameMode `json:"modes"`
}
