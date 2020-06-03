package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB connection holder
type MongoDB struct {
	client   *mongo.Client
	database string
}

// NewDB creates a DB connection and returns a db instance
func NewDB(ctx context.Context, uri, database string) (db *MongoDB, err error) {
	db = &MongoDB{}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	db.database = database
	db.client = client
	return
}

// Disconnect closes the mongodb connection
func (db *MongoDB) Disconnect(ctx context.Context) {
	db.client.Disconnect(ctx)
}

// Ping db
func (db *MongoDB) Ping(ctx context.Context) (bool, error) {
	err := db.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetCollection func
func (db *MongoDB) GetCollection(collection string) *mongo.Collection {
	return db.client.Database(db.database).Collection(collection)
}
