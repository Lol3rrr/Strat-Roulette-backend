package database

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func (s *session) loadCredsAndReconnect() error {
	data, err := s.VaultSession.ReadData("database/creds/strat-roulette")
	if err != nil {
		return errors.New("Could not load Credentials from Vault")
	}

	creds := data.Data
	username, found := creds["username"]
	if !found {
		return errors.New("Vault response did not include a username")
	}
	password, found := creds["password"]
	if !found {
		return errors.New("Vault response did not include a password")
	}

	s.Username = username.(string)
	s.Password = password.(string)

	s.VaultSession.RenewDataForever(data, data.LeaseDuration, func() {
		s.loadCredsAndReconnect()

		err = s.connect()
		if err != nil {
			logrus.Errorf("Connecting to Database: %v \n", err)
		}
	})

	return nil
}
