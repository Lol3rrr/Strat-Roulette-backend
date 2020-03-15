package database

import "go.mongodb.org/mongo-driver/mongo"

type session struct {
	URL             string
	Port            string
	Database        string
	Collection      string
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
}
