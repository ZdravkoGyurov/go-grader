package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db/models"
)

// DBHandler ...
type DBHandler struct {
	appCtx     app.Context
	collection *mongo.Collection
}

// NewHandler creates a new user DB handler
func NewDBHandler(appCtx app.Context, client *mongo.Client) *DBHandler {
	return &DBHandler{
		collection: client.Database(appCtx.Cfg.DatabaseName).
			Collection(models.UserCollectionName),
	}
}

// Create creates new user in the database
func (h *DBHandler) Create(ctx context.Context, user *models.User) error {
	_, err := h.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// Read ...
func (h *DBHandler) ReadByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User

	result := h.collection.FindOne(ctx, bson.M{"_id": userID})
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user with id %s: %w", userID, err)
	}

	return &user, nil
}

// Read ...
func (h *DBHandler) ReadByUsername(ctx context.Context, username string) (*models.User, error) {
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
