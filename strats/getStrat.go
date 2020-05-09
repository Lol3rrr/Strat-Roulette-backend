package strats

import (
	"errors"

	"github.com/Lol3rrr/mongovault"
)

func (s *session) GetStrat(id string) (Strat, error) {
	if len(id) <= 0 {
		return Strat{}, errors.New("ID can not be empty")
	}

	query := []mongovault.Filter{
		{
			Key:   "id",
			Value: id,
		},
	}
	var result Strat
	err := s.Database.Get(query, &result)

	return result, err
}
