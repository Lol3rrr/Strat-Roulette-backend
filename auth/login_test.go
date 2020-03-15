package auth

import (
	"errors"
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	tables := []struct {
		Name          string
		InputSession  session
		InputDBError  error
		InputUsername string
		InputPassword string
		ResultError   bool
	}{
		{
			Name: "Valid Input",
			InputSession: session{
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
			},
			InputDBError:  nil,
			InputUsername: "testAdmin",
			InputPassword: "testPassword",
			ResultError:   false,
		},
		{
			Name: "Database returns error",
			InputSession: session{
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
			},
			InputDBError:  errors.New("testError"),
			InputUsername: "testAdmin",
			InputPassword: "testPassword",
			ResultError:   true,
		},
		{
			Name: "Invalid Input, empty username",
			InputSession: session{
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
			},
			InputDBError:  nil,
			InputUsername: "",
			InputPassword: "testPassword",
			ResultError:   true,
		},
		{
			Name: "Invalid Input, empty password",
			InputSession: session{
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
			},
			InputDBError:  nil,
			InputUsername: "testAdmin",
			InputPassword: "",
			ResultError:   true,
		},
		{
			Name: "Invalid Input, mismatching usernames",
			InputSession: session{
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
			},
			InputDBError:  nil,
			InputUsername: "testName",
			InputPassword: "testPassword",
			ResultError:   true,
		},
		{
			Name: "Invalid Input, mismatching passwords",
			InputSession: session{
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
			},
			InputDBError:  nil,
			InputUsername: "testAdmin",
			InputPassword: "somePassword",
			ResultError:   true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBError := table.InputDBError
		inUsername := table.InputUsername
		inPassword := table.InputPassword
		resultError := table.ResultError

		inSession.Database = &database.MockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "Insert",
						ReturnArguments: mock.Arguments{
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			_, outputError := inSession.Login(inUsername, inPassword)

			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
