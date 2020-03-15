package database

import (
	"context"
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

	_, err := s.MongoCollection.DeleteOne(ctx, filter)

	return err
}
