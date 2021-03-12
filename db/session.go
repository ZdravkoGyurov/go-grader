package db

import (
	"context"
	"fmt"
	"grader/app"
	"grader/db/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SessionHandler ...
type SessionHandler struct {
	appCtx     app.Context
	collection *mongo.Collection
}

// NewSessionHandler creates a new session DB handler
func NewSessionHandler(appCtx app.Context, client *mongo.Client) *SessionHandler {
	return &SessionHandler{
		collection: client.Database(appCtx.Cfg.DatabaseName).
			Collection(models.SessionCollectionName),
	}
}

// CreateSession creates new session in the database
func (h *SessionHandler) CreateSession(ctx context.Context, session *models.Session) error {
	_, err := h.collection.InsertOne(ctx, session)
	if err != nil {
		return fmt.Errorf("failed to insert session: %w", err)
	}

	return nil
}

// ReadSession ...
func (h *SessionHandler) ReadSession(ctx context.Context, sessionID string) (*models.Session, error) {
	var session models.Session
	if err := h.collection.FindOne(ctx, filterByID(sessionID)).Decode(&session); err != nil {
		return nil, fmt.Errorf("failed to find session with id %s: %w", sessionID, err)
	}

	return &session, nil
}

// DeleteSession ...
func (h *SessionHandler) DeleteSession(ctx context.Context, sessionID string) error {
	if _, err := h.collection.DeleteOne(ctx, filterByID(sessionID)); err != nil {
		return fmt.Errorf("failed to delete session with id %s: %w", sessionID, err)
	}

	return nil
}
