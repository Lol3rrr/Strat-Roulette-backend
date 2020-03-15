package database

import "go.mongodb.org/mongo-driver/bson/primitive"

func convertQueryToPrimitive(query map[string]interface{}) []primitive.E {
	result := make([]primitive.E, len(query))

	current := 0
	for key, value := range query {
		result[current] = primitive.E{
			Key:   key,
			Value: value,
		}

		current++
	}

	return result
}
