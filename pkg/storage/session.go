package storage

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const sessionCollection = "sessions"

func (s *Storage) CreateSession(ctx context.Context, session *model.Session) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(sessionCollection)
	_, err := collection.InsertOne(ctx, session)
	if err != nil {
		return errors.Wrap(err, "failed to insert session")
	}

	return nil
}

func (s *Storage) ReadSession(ctx context.Context, sessionID string) (*model.Session, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(sessionCollection)
	var session model.Session
	if err := collection.FindOne(ctx, filterByID(sessionID)).Decode(&session); err != nil {
		return nil, errors.Wrapf(err, "failed to find session with id %s", sessionID)
	}

	return &session, nil
}

func (s *Storage) DeleteSession(ctx context.Context, sessionID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(sessionCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(sessionID)); err != nil {
		return errors.Wrapf(err, "failed to delete session with id %s", err)
	}

	return nil
}
