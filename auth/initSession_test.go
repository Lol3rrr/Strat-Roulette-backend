package auth

import (
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitSession(t *testing.T) {
	tables := []struct{
		Name string
		InputDBSession database.SessionInterface
		InputAdminUsername string
		InputAdminPassword string
		InputSessionDuration int64
		Result SessionInterface
	}{
		{
			Name: "Valid Input",
			InputDBSession: &database.MockDatabase{},
			InputAdminUsername: "testAdmin",
			InputAdminPassword: "testPassword",
			InputSessionDuration: 123,
			Result: &session{
				Database: &database.MockDatabase{},
				AdminUsername: "testAdmin",
				AdminPassword: "testPassword",
				SessionDuration: 123 * 60,
			},
		},
	}

	for _, table := range tables {
		inDBSession := table.InputDBSession
		inAdminUsername := table.InputAdminUsername
		inAdminPassword := table.InputAdminPassword
		inSessionDuration := table.InputSessionDuration
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := InitSession(inDBSession, inAdminUsername, inAdminPassword, inSessionDuration)

			assert.Equal(t, result, output)
		})
	}
}