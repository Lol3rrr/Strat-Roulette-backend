package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	tables := []struct {
		Name          string
		EnvKey        string
		EnvValue      string
		InputKey      string
		InputFallback string
		Result        string
	}{
		{
			Name:          "Valid Input",
			EnvKey:        "testKey",
			EnvValue:      "testValue",
			InputKey:      "testKey",
			InputFallback: "testFallback",
			Result:        "testValue",
		},
		{
			Name:          "Key not found",
			EnvKey:        "testKey",
			EnvValue:      "testValue",
			InputKey:      "someKey",
			InputFallback: "testFallback",
			Result:        "testFallback",
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

			output := GetEnvString(inputKey, inputFallback)

			assert.Equal(t, result, output)
		})
	}
}
