package strats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMode(t *testing.T) {
	tables := []struct {
		Name      string
		InputMode GameMode
		Result    bool
	}{
		{
			Name:      "Valid Input, 'Bomb'",
			InputMode: Bomb,
			Result:    true,
		},
		{
			Name:      "Valid Input, 'Secure Area'",
			InputMode: SecureArea,
			Result:    true,
		},
		{
			Name:      "Valid Input, 'Hostage'",
			InputMode: Hostage,
			Result:    true,
		},
		{
			Name:      "Invalid Input",
			InputMode: "testMode",
			Result:    false,
		},
	}

	for _, table := range tables {
		inMode := table.InputMode
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := isValidMode(inMode)

			assert.Equal(t, result, output)
		})
	}
}
