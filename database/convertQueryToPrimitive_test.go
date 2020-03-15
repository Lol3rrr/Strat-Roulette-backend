package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestConvertQueryToPrimitive(t *testing.T) {
	tables := []struct {
		Name       string
		InputQuery map[string]interface{}
		Result     []primitive.E
	}{
		{
			Name: "Valid Input",
			InputQuery: map[string]interface{}{
				"testKey":  "testValue",
				"testKey2": "testValue2",
			},
			Result: []primitive.E{
				{
					Key:   "testKey",
					Value: "testValue",
				},
				{
					Key:   "testKey2",
					Value: "testValue2",
				},
			},
		},
	}

	for _, table := range tables {
		inQuery := table.InputQuery
		result := table.Result

		t.Run(table.Name, func(t *testing.T) {
			output := convertQueryToPrimitive(inQuery)

			assert.Equal(t, result, output)
		})
	}
}
