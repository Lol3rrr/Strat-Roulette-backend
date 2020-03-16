package auth

import (
	"errors"
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCleanUpSessions(t *testing.T) {
	tables := []struct {
		Name         string
		InputSession session
		InputDBError error
		InputTime    int64
		ResultError  bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputDBError: nil,
			InputTime:    int64(123),
			ResultError:  false,
		},
		{
			Name:         "Database returns error",
			InputSession: session{},
			InputDBError: errors.New("testError"),
			InputTime:    int64(123),
			ResultError:  true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBError := table.InputDBError
		inTime := table.InputTime
		resultError := table.ResultError

		inSession.Database = &database.MockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "DeleteMany",
						ReturnArguments: mock.Arguments{
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			outputError := inSession.CleanUpSessions(inTime)

			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
