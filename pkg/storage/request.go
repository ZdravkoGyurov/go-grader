package storage

import (
	"context"
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

const requestCollection = "requests"

func (s *Storage) CreateRequest(ctx context.Context, request *model.Request) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(requestCollection)
	_, err := collection.InsertOne(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to insert request: %w", err)
	}

	return nil
}

func (s *Storage) ReadRequest(ctx context.Context, requestID string) (*model.Request, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(requestCollection)
	var request model.Request
	if err := collection.FindOne(ctx, filterByID(requestID)).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to find request with id %s: %w", requestID, err)
	}

	return &request, nil
}

// ReadAllRequests reads all requests with optional userID filter
func (s *Storage) ReadAllRequests(ctx context.Context, userID string) ([]*model.Request, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(requestCollection)

	cursor, err := collection.Find(ctx, filterRequestByUserID(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to find all requests with user_id %s: %w", userID, err)
	}

	var requests []*model.Request
	if err = cursor.All(ctx, &requests); err != nil {
		return nil, fmt.Errorf("failed to decode all requests with user_id %s: %w", userID, err)
	}

	return requests, nil
}
func (s *Storage) UpdateRequest(ctx context.Context, requestID string, request *model.Request) (*model.Request, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(requestCollection)
	var updatedRequest model.Request
	result := collection.FindOneAndUpdate(ctx, filterByID(requestID), update(request), updateOpts())
	if err := result.Decode(&updatedRequest); err != nil {
		return nil, fmt.Errorf("failed to find and update request with id %s: %w", requestID, err)
	}

	return &updatedRequest, nil
}

func (s *Storage) DeleteRequest(ctx context.Context, requestID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(requestCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(requestID)); err != nil {
		return fmt.Errorf("failed to delete request with id %s: %w", requestID, err)
	}
	return nil
}
