package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (s *session) isConnectionAlive() bool {
	if s.MongoClient == nil {
		return false
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()

	err := s.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return false
	}

	return true
}
