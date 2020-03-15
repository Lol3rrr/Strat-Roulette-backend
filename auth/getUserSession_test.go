package auth

import (
	"errors"
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserSession(t *testing.T) {
	tables := []struct {
		Name              string
		InputSession      session
		InputDBValue      userSession
		InputDBError      error
		InputSessionID    string
		ResultUserSession UserSessionInterface
		ResultError       bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputDBValue: userSession{
				SessionID:  "testSessionID",
				UserRole:   Admin,
				Created:    123,
				Expiration: 234,
			},
			InputDBError:   nil,
			InputSessionID: "testSessionID",
			ResultUserSession: &userSession{
				SessionID:  "testSessionID",
				UserRole:   Admin,
				Created:    123,
				Expiration: 234,
			},
			ResultError: false,
		},
		{
			Name:              "Database returns error",
			InputSession:      session{},
			InputDBValue:      userSession{},
			InputDBError:      errors.New("testError"),
			InputSessionID:    "testSessionID",
			ResultUserSession: &userSession{},
			ResultError:       true,
		},
		{
			Name:         "Invalid Input, empty SessionID",
			InputSession: session{},
			InputDBValue: userSession{
				SessionID:  "testSessionID",
				UserRole:   Admin,
				Created:    123,
				Expiration: 234,
			},
			InputDBError:      nil,
			InputSessionID:    "",
			ResultUserSession: &userSession{},
			ResultError:       true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBValue := table.InputDBValue
		inDBError := table.InputDBError
		inSessionID := table.InputSessionID
		resultUserSession := table.ResultUserSession
		resultError := table.ResultError

		inSession.Database = &database.MockDatabase{
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
			outputUserSession, outputError := inSession.GetUserSession(inSessionID)

			assert.Equal(t, resultUserSession, outputUserSession)
			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
