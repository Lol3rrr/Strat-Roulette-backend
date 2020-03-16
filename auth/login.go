package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

func (s *session) Login(username, password string) (UserSessionInterface, error) {
	if username != s.AdminUsername {
		return &userSession{}, errors.New("Unknown username")
	}
	if password != s.AdminPassword {
		return &userSession{}, errors.New("Unknown password")
	}

	now := time.Now().Unix()
	expiration := now + s.SessionDuration

	result := &userSession{
		SessionID:  uuid.New().String(),
		UserRole:   Admin,
		Created:    now,
		Expiration: expiration,
	}

	err := s.Database.Insert(result)

	return result, err
}
