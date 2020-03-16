package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (s *session) connect() error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+s.URL+":"+s.Port))
	if err != nil {
		return err
	}

	pingCtx, cancelPingCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelPingCtx()

	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		return err
	}

	s.MongoClient = client
	s.MongoCollection = client.Database(s.Database).Collection(s.Collection)

	return nil
}
