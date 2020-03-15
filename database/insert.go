package database

import (
	"context"
	"time"
)

// Insert adds the given interface to the database
func (s *session) Insert(data interface{}) error {
	if !s.isConnectionAlive() {
		err := s.Connect()
		if err != nil {
			return err
		}
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	_, err := s.MongoCollection.InsertOne(ctx, data)

	return err
}
