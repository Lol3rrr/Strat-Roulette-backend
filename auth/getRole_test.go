package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRole(t *testing.T) {
	tables := []struct{
		Name string
		InputUserSession userSession
		Result Role
	}{
		{
			Name: "Valid Input",
			InputUserSession: userSession{
				UserRole: Admin,
			},
			Result: Admin,
		},
	}

	for _, table := range tables {
		inUserSession := table.InputUserSession
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := inUserSession.GetRole()

			assert.Equal(t, result, output)
		})
	}
}