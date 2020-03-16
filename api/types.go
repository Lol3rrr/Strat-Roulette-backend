package api

import (
	"strat-roulette-backend/auth"
	"strat-roulette-backend/strats"
)

type session struct {
	Strats strats.SessionInterface
	Auth   auth.SessionInterface
}

type addStratRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Site        strats.Site       `json:"site"`
	Modes       []strats.GameMode `json:"modes"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type allStratsResponse struct {
	Strats []strats.Strat `json:"strats"`
}
