package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExpiration(t *testing.T) {
	tables := []struct {
		Name             string
		InputUserSession userSession
		Result           int64
	}{
		{
			Name: "Valid Input",
			InputUserSession: userSession{
				Expiration: 123,
			},
			Result: 123,
		},
	}

	for _, table := range tables {
		inUserSession := table.InputUserSession
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := inUserSession.GetExpiration()

			assert.Equal(t, result, output)
		})
	}
}
