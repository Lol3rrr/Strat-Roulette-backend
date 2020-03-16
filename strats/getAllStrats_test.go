package strats

import (
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllStrats(t *testing.T) {
	tables := []struct {
		Name         string
		InputSession session
		InputDBValue []Strat
		InputDBError error
		ResultStrats []Strat
		ResultError  bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputDBValue: []Strat{
				{
					ID:          "testID1",
					Name:        "testName1",
					Description: "testDescription1",
					PlayerSite:  Attacker,
					Modes: []GameMode{
						Bomb,
					},
				},
				{
					ID:          "testID2",
					Name:        "testName2",
					Description: "testDescription2",
					PlayerSite:  Defender,
					Modes: []GameMode{
						Bomb,
						Hostage,
					},
				},
			},
			InputDBError: nil,
			ResultStrats: []Strat{
				{
					ID:          "testID1",
					Name:        "testName1",
					Description: "testDescription1",
					PlayerSite:  Attacker,
					Modes: []GameMode{
						Bomb,
					},
				},
				{
					ID:          "testID2",
					Name:        "testName2",
					Description: "testDescription2",
					PlayerSite:  Defender,
					Modes: []GameMode{
						Bomb,
						Hostage,
					},
				},
			},
			ResultError: false,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBValue := table.InputDBValue
		inDBError := table.InputDBError
		resultStrats := table.ResultStrats
		resultError := table.ResultError

		inSession.Database = &database.MockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetAll",
						ReturnArguments: mock.Arguments{
							inDBValue,
							len(inDBValue),
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			outputStrats, outputError := inSession.GetAllStrats()

			assert.Equal(t, resultStrats, outputStrats)
			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
