package strats

import (
	"strat-roulette-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitSession(t *testing.T) {
	tables := []struct {
		Name           string
		InputDBSession database.SessionInterface
		Result         SessionInterface
	}{
		{
			Name:           "Valid Input",
			InputDBSession: &database.MockDatabase{},
			Result: &session{
				Database: &database.MockDatabase{},
			},
		},
	}

	for _, table := range tables {
		inDBSession := table.InputDBSession
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := InitSession(inDBSession)

			assert.Equal(t, result, output)
		})
	}
}
