package strats

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRandomStrat(t *testing.T) {
	tables := []struct {
		Name          string
		InputSession  session
		InputDBResult []Strat
		InputDBError  error
		InputSite     Site
		InputMode     GameMode
		ResultError   bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputDBResult: []Strat{
				{
					ID:          "testID",
					Name:        "testName",
					Description: "testDescription",
					Modes: []GameMode{
						Bomb,
						SecureArea,
					},
					PlayerSite: Attacker,
				},
			},
			InputDBError: nil,
			InputSite:    Attacker,
			InputMode:    Bomb,
			ResultError:  false,
		},
		{
			Name:         "Database returns error",
			InputSession: session{},
			InputDBResult: []Strat{
				{
					ID:          "testID",
					Name:        "testName",
					Description: "testDescription",
					Modes: []GameMode{
						Bomb,
						SecureArea,
					},
					PlayerSite: Attacker,
				},
			},
			InputDBError: errors.New("testError"),
			InputSite:    Attacker,
			InputMode:    Bomb,
			ResultError:  true,
		},
		{
			Name:          "Database returns empty array",
			InputSession:  session{},
			InputDBResult: []Strat{},
			InputDBError:  nil,
			InputSite:     Attacker,
			InputMode:     Bomb,
			ResultError:   true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBResult := table.InputDBResult
		inDBError := table.InputDBError
		inSite := table.InputSite
		inMode := table.InputMode
		resultError := table.ResultError

		inSession.Database = &mockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetAll",
						ReturnArguments: mock.Arguments{
							inDBResult,
							len(inDBResult),
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			outputStrat, outputError := inSession.GetRandomStrat(inSite, inMode)

			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
				assert.Equal(t, inSite, outputStrat.PlayerSite)
				assert.Contains(t, outputStrat.Modes, inMode)
			}
		})
	}
}
