package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Update updates the first entry that matches the query with the given Data
func (s *session) Update(query map[string]interface{}, updateData map[string]interface{}) error {
	if !s.isConnectionAlive() {
		err := s.Connect()
		if err != nil {
			return err
		}
	}

	update := bson.M{
		"$set": updateData,
	}

	filter := convertQueryToPrimitive(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := s.MongoCollection.UpdateOne(ctx, filter, update)

	return err
}
