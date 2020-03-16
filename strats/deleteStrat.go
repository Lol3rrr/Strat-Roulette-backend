package strats

import "errors"

func (s *session) DeleteStrat(id string) error {
	if len(id) <= 0 {
		return errors.New("ID can not be empty")
	}

	query := map[string]interface{}{
		"id": id,
	}
	return s.Database.Delete(query)
}
