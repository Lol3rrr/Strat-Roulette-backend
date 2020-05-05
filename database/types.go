package database

import (
	"github.com/Lol3rrr/cvault"
	"go.mongodb.org/mongo-driver/mongo"
)

type session struct {
	URL             string
	Port            string
	Database        string
	Collection      string
	Username        string
	Password        string
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
	VaultSession    cvault.Session
}
