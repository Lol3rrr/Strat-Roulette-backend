package strats

import (
	"errors"
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteStrat(t *testing.T) {
	tables := []struct {
		Name         string
		InputSession session
		InputDBError error
		InputID      string
		ResultError  bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputDBError: nil,
			InputID:      "testID",
			ResultError:  false,
		},
		{
			Name:         "Database returns error",
			InputSession: session{},
			InputDBError: errors.New("testError"),
			InputID:      "testID",
			ResultError:  true,
		},
		{
			Name:         "Invalid Input, empty ID",
			InputSession: session{},
			InputDBError: nil,
			InputID:      "",
			ResultError:  true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBError := table.InputDBError
		inID := table.InputID
		resultError := table.ResultError

		inSession.Database = &database.MockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "Delete",
						ReturnArguments: mock.Arguments{
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			outputError := inSession.DeleteStrat(inID)

			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
