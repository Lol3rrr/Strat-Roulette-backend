package strats

import "errors"

func (s *session) GetStrat(id string) (Strat, error) {
	if len(id) <= 0 {
		return Strat{}, errors.New("ID can not be empty")
	}

	query := map[string]interface{}{
		"id": id,
	}
	var result Strat
	err := s.Database.Get(query, &result)

	return result, err
}
