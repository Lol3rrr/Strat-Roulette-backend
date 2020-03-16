package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvInt(t *testing.T) {
	tables := []struct {
		Name          string
		EnvKey        string
		EnvValue      string
		InputKey      string
		InputFallback int
		Result        int
	}{
		{
			Name:          "Valid Input",
			EnvKey:        "testKey",
			EnvValue:      "1",
			InputKey:      "testKey",
			InputFallback: 3,
			Result:        1,
		},
		{
			Name:          "Key not found",
			EnvKey:        "testKey",
			EnvValue:      "1",
			InputKey:      "someKey",
			InputFallback: 3,
			Result:        3,
		},
		{
			Name:          "Value is not a number",
			EnvKey:        "testKey",
			EnvValue:      "testValue",
			InputKey:      "testKey",
			InputFallback: 3,
			Result:        3,
		},
	}

	for _, table := range tables {
		envKey := table.EnvKey
		envValue := table.EnvValue
		inputKey := table.InputKey
		inputFallback := table.InputFallback
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			err := os.Setenv(envKey, envValue)
			if err != nil {
				t.Fatal(err)
			}

			output := GetEnvInt(inputKey, inputFallback)

			assert.Equal(t, result, output)
		})
	}
}
