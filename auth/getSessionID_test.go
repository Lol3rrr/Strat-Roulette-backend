package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSessionID(t *testing.T) {
	tables := []struct {
		Name             string
		InputUserSession userSession
		Result           string
	}{
		{
			Name: "Valid Input",
			InputUserSession: userSession{
				SessionID: "testSessionID",
			},
			Result: "testSessionID",
		},
	}

	for _, table := range tables {
		inUserSession := table.InputUserSession
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := inUserSession.GetSessionID()

			assert.Equal(t, result, output)
		})
	}
}
