package session

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

// NewDBHandler creates a new session DB handler
func NewDBHandler(appCtx app.Context, client *mongo.Client) *DBHandler {
	return &DBHandler{
		collection: client.Database(appCtx.Cfg.DatabaseName).
			Collection(models.SessionCollectionName),
	}
}

// Create creates new session in the database
func (h *DBHandler) Create(ctx context.Context, session *models.Session) error {
	_, err := h.collection.InsertOne(ctx, session)
	if err != nil {
		return fmt.Errorf("failed to insert session: %w", err)
	}

	return nil
}

// Read ...
func (h *DBHandler) Read(ctx context.Context, sessionID string) (*models.Session, error) {
	var session models.Session
	if err := h.collection.FindOne(ctx, bson.M{"_id": sessionID}).Decode(&session); err != nil {
		return nil, fmt.Errorf("failed to find session with id %s: %w", sessionID, err)
	}

	return &session, nil
}

// Delete ...
func (h *DBHandler) Delete(ctx context.Context, sessionID string) error {
	if _, err := h.collection.DeleteOne(ctx, bson.M{"_id": sessionID}); err != nil {
		return fmt.Errorf("failed to delete session with id %s: %w", sessionID, err)
	}

	return nil
}
