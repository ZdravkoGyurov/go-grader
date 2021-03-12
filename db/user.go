package db

import (
	"context"
	"fmt"
	"grader/app"
	"grader/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserHandler ...
type UserHandler struct {
	appCtx     app.Context
	collection *mongo.Collection
}

// NewUserHandler creates a new user DB handler
func NewUserHandler(appCtx app.Context, client *mongo.Client) *UserHandler {
	return &UserHandler{
		collection: client.Database(appCtx.Cfg.DatabaseName).
			Collection(models.UserCollectionName),
	}
}

// CreateUser creates new user in the database
func (h *UserHandler) CreateUser(ctx context.Context, user *models.User) error {
	_, err := h.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// ReadUser ...
func (h *UserHandler) ReadUser(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	result := h.collection.FindOne(ctx, filterByUsername(username))
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user with username %s: %w", username, err)
	}

	return &user, nil
}

func filterByUsername(username string) bson.M {
	return bson.M{"username": username}
}
