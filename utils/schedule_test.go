package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	tables := []struct {
		Name          string
		InputDuration time.Duration
		InputWait     bool
		ResultRuns    int
	}{
		{
			Name:          "Valid Input",
			InputDuration: time.Second * 1,
			InputWait:     true,
			ResultRuns:    2,
		},
		{
			Name:          "Immediate stop",
			InputDuration: time.Second * 1,
			InputWait:     false,
			ResultRuns:    1,
		},
	}

	for _, table := range tables {
		inDuration := table.InputDuration
		inWait := table.InputWait
		resultRuns := table.ResultRuns

		t.Run(table.Name, func(t *testing.T) {
			t.Parallel()

			outputRuns := 0

			stopChannel := Schedule(func() {
				outputRuns++
			}, inDuration)

			if inWait {
				time.Sleep(inDuration)
			}

			stopChannel <- true

			assert.Equal(t, resultRuns, outputRuns)
		})
	}
}
