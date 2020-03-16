package strats

// GetAllStrats returns all
func (s *session) GetAllStrats() ([]Strat, error) {
	var result []Strat
	err := s.Database.GetAll(map[string]interface{}{}, &result)

	return result, err
}
