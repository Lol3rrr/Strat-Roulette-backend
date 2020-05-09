package strats

import "github.com/Lol3rrr/mongovault"

// GetAllStrats returns all
func (s *session) GetAllStrats() ([]Strat, error) {
	var result []Strat
	err := s.Database.GetAll([]mongovault.Filter{}, &result)

	return result, err
}
