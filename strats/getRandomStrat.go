package strats

import (
	"errors"
	"math/rand"

	"github.com/Lol3rrr/mongovault"
)

// GetRandomStrat selects one random Strat that matches the given Params
func (s *session) GetRandomStrat(playerSite Site, mode GameMode) (Strat, error) {
	if !isValidSite(playerSite) {
		return Strat{}, errors.New("Unknown PlayerSite")
	}
	if !isValidMode(mode) {
		return Strat{}, errors.New("Unknown Mode")
	}

	var entrys []Strat
	query := []mongovault.Filter{
		{
			Key:   "playerSite",
			Value: playerSite,
		},
		{
			Key:   "modes",
			Value: mode,
		},
	}
	err := s.Database.GetAll(query, &entrys)
	if err != nil {
		return Strat{}, err
	}

	if len(entrys) <= 0 {
		return Strat{}, errors.New("No Strat found that matches the Params")
	}

	return entrys[rand.Intn(len(entrys))], nil
}
