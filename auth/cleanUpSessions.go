package auth

import "github.com/Lol3rrr/mongovault"

func (s *session) CleanUpSessions(now int64) error {
	query := []mongovault.Filter{
		{
			Key: "expiration",
			Value: mongovault.Filter{
				Key:   "$lt",
				Value: now,
			},
		},
	}

	return s.Database.DeleteMany(query)
}
