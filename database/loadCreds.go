package database

import (
	"errors"

	"github.com/hashicorp/vault/api"
)

func (s *session) loadCreds() (*api.Secret, error) {
	data, err := s.VaultSession.ReadData("database/creds/strat-roulette")
	if err != nil {
		return nil, errors.New("Could not load Credentials from Vault")
	}

	creds := data.Data
	username, found := creds["username"]
	if !found {
		return nil, errors.New("Vault response did not include a username")
	}
	password, found := creds["password"]
	if !found {
		return nil, errors.New("Vault response did not include a password")
	}

	s.Username = username.(string)
	s.Password = password.(string)

	return data, nil
}
