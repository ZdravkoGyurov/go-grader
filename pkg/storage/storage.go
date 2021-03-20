package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ZdravkoGyurov/go-grader/pkg/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type Storage struct {
	config      config.Config
	mongoClient *mongo.Client
}

func New(ctx context.Context, config config.Config) (*Storage, error) {
	mongoClient, err := connect(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}

	return &Storage{
		config:      config,
		mongoClient: mongoClient,
	}, nil
}

func (s *Storage) Disconnect() {
	dbDisconnectCtx, cancel := context.WithTimeout(context.Background(), s.config.DBDisconnectTimeout)
	defer cancel()
	err := s.mongoClient.Disconnect(dbDisconnectCtx)
	if err != nil {
		log.Error().Printf("failed to disconnect from db: %s\n", err)
	}
}

func connect(ctx context.Context, config config.Config) (*mongo.Client, error) {
	options := options.Client().ApplyURI(config.DatabaseURI)

	ctx, cancel := context.WithTimeout(ctx, config.DBConnectTimeout)
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
