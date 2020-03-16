package strats

import (
	"errors"

	"github.com/google/uuid"
)

// AddStrat adds a new strat with the given Parameters
func (s *session) AddStrat(pName, pDescription string, pSite Site, pModes []GameMode) error {
	if len(pName) <= 0 {
		return errors.New("Name was empty")
	}
	if len(pDescription) <= 0 {
		return errors.New("Description was empty")
	}
	if len(pSite) <= 0 {
		return errors.New("Site was empty")
	}
	if len(pModes) <= 0 {
		return errors.New("Modes was empty")
	}

	if !isValidSite(pSite) {
		return errors.New("Unknown PlayerSite")
	}

	for _, mode := range pModes {
		if !isValidMode(mode) {
			return errors.New("Unknown Mode")
		}
	}

	id := uuid.New().String()

	nStrat := Strat{
		ID:          id,
		Name:        pName,
		Description: pDescription,
		PlayerSite:  pSite,
		Modes:       pModes,
	}

	return s.Database.Insert(nStrat)
}
