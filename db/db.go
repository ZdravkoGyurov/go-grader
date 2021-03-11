package db

import (
	"context"
	"grader/app/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect ...
func Connect(config config.Config) (*mongo.Client, error) {
	options := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), config.DBConnectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
