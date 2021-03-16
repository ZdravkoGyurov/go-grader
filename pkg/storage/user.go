package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const userCollection = "users"

func (s *Storage) CreateUser(ctx context.Context, user *model.User) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (s *Storage) ReadUserByID(ctx context.Context, userID string) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var user model.User

	result := collection.FindOne(ctx, bson.M{"_id": userID})
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user with id %s: %w", userID, err)
	}

	return &user, nil
}

func (s *Storage) ReadUserByUsername(ctx context.Context, username string) (*model.User, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(userCollection)
	var user model.User

	result := collection.FindOne(ctx, filterByUsername(username))
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user with username %s: %w", username, err)
	}

	return &user, nil
}

func filterByUsername(username string) bson.M {
	return bson.M{"username": username}
}
