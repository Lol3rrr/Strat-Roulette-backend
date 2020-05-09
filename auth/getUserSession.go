package auth

import (
	"errors"

	"github.com/Lol3rrr/mongovault"
)

// GetUserSession simply loads the UserSession for the given SessionID
func (s *session) GetUserSession(sessionID string) (UserSessionInterface, error) {
	if len(sessionID) <= 0 {
		return &userSession{}, errors.New("Session ID can not be empty")
	}

	query := []mongovault.Filter{
		{
			Key:   "sessionID",
			Value: sessionID,
		},
	}
	var result userSession
	err := s.Database.Get(query, &result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
