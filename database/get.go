package database

import (
	"context"
	"time"
)

// Get is used to load a single entry from the Database
func (s *session) Get(query map[string]interface{}, result interface{}) error {
	if !s.isConnectionAlive() {
		err := s.connect()
		if err != nil {
			return err
		}
	}

	filter := convertQueryToPrimitive(query)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	err := s.MongoCollection.FindOne(ctx, filter).Decode(result)

	return err
}
