package database

import (
	"context"
	"errors"
	"time"
)

// Delete deletes the first entry matching the query
func (s *session) Delete(query map[string]interface{}) error {
	if !s.isConnectionAlive() {
		err := s.connect()
		if err != nil {
			return err
		}
	}

	filter := convertQueryToPrimitive(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	result, err := s.MongoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount <= 0 {
		return errors.New("Could not find any element matching the Query")
	}

	return nil
}
