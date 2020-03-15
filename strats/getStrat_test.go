package strats

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetStrat(t *testing.T) {
	tables := []struct {
		Name         string
		InputSession session
		InputDBValue Strat
		InputDBError error
		InputID      string
		ResultStrat  Strat
		ResultError  bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputDBValue: Strat{
				ID:          "testID",
				Name:        "testName",
				Description: "testDescription",
				PlayerSite:  Attacker,
				Modes: []GameMode{
					Bomb,
				},
			},
			InputDBError: nil,
			InputID:      "testID",
			ResultStrat: Strat{
				ID:          "testID",
				Name:        "testName",
				Description: "testDescription",
				PlayerSite:  Attacker,
				Modes: []GameMode{
					Bomb,
				},
			},
			ResultError: false,
		},
		{
			Name:         "Database returns error",
			InputSession: session{},
			InputDBValue: Strat{},
			InputDBError: errors.New("testError"),
			InputID:      "testID",
			ResultStrat:  Strat{},
			ResultError:  true,
		},
		{
			Name:         "Empty ID",
			InputSession: session{},
			InputDBValue: Strat{
				ID:          "testID",
				Name:        "testName",
				Description: "testDescription",
				PlayerSite:  Attacker,
				Modes: []GameMode{
					Bomb,
				},
			},
			InputDBError: nil,
			InputID:      "",
			ResultStrat:  Strat{},
			ResultError:  true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBValue := table.InputDBValue
		inDBError := table.InputDBError
		inID := table.InputID
		resultStrat := table.ResultStrat
		resultError := table.ResultError

		inSession.Database = &mockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "Get",
						ReturnArguments: mock.Arguments{
							inDBValue,
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			outputStrat, outputError := inSession.GetStrat(inID)

			assert.Equal(t, resultStrat, outputStrat)
			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
