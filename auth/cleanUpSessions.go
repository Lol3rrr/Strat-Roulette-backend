package auth

func (s *session) CleanUpSessions(now int64) error {
	query := map[string]interface{}{
		"expiration": map[string]interface{}{
			"$lt": now,
		},
	}
	return s.Database.DeleteMany(query)
}
