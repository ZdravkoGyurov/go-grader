package storage

import (
	"context"
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const sessionCollection = "sessions"

func (s *Storage) CreateSession(ctx context.Context, session *model.Session) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(sessionCollection)
	_, err := collection.InsertOne(ctx, session)
	if err != nil {
		return fmt.Errorf("failed to insert session: %w", err)
	}

	return nil
}

func (s *Storage) ReadSession(ctx context.Context, sessionID string) (*model.Session, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(sessionCollection)
	var session model.Session
	if err := collection.FindOne(ctx, filterByID(sessionID)).Decode(&session); err != nil {
		return nil, fmt.Errorf("failed to find session with id %s: %w", sessionID, err)
	}

	return &session, nil
}

func (s *Storage) DeleteSession(ctx context.Context, sessionID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(sessionCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(sessionID)); err != nil {
		return fmt.Errorf("failed to delete session with id %s: %w", sessionID, err)
	}

	return nil
}
