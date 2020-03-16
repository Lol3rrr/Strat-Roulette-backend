package strats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSite(t *testing.T) {
	tables := []struct {
		Name      string
		InputSite Site
		Result    bool
	}{
		{
			Name:      "Valid Input, 'Attacker'",
			InputSite: Attacker,
			Result:    true,
		},
		{
			Name:      "Valid Input, 'Defender'",
			InputSite: Defender,
			Result:    true,
		},
		{
			Name:      "Invalid Input",
			InputSite: "testSite",
			Result:    false,
		},
	}

	for _, table := range tables {
		inSite := table.InputSite
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := isValidSite(inSite)

			assert.Equal(t, result, output)
		})
	}
}
