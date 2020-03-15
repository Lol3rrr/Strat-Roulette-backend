package strats

import (
	"errors"
	"math/rand"
)

// GetRandomStrat selects one random Strat that matches the given Params
func (s *session) GetRandomStrat(playerSite Site, mode GameMode) (Strat, error) {
	var entrys []Strat
	query := map[string]interface{}{
		"playerSite": playerSite,
		"modes":      mode,
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
