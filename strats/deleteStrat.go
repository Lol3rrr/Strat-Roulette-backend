package strats

import (
	"errors"

	"github.com/Lol3rrr/mongovault"
)

func (s *session) DeleteStrat(id string) error {
	if len(id) <= 0 {
		return errors.New("ID can not be empty")
	}

	query := []mongovault.Filter{
		{
			Key:   "id",
			Value: id,
		},
	}
	return s.Database.Delete(query)
}
