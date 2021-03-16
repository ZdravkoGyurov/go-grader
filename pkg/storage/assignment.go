package storage

import (
	"context"
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const assignmentCollection = "assignments"

func (s *Storage) CreateAssignment(ctx context.Context, assignment *model.Assignment) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)
	_, err := collection.InsertOne(ctx, assignment)
	if err != nil {
		return fmt.Errorf("failed to insert assignment: %w", err)
	}

	return nil
}

func (s *Storage) ReadAssignment(ctx context.Context, assignmentID string) (*model.Assignment, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)
	var assignment model.Assignment
	if err := collection.FindOne(ctx, bson.M{"_id": assignmentID}).Decode(&assignment); err != nil {
		return nil, fmt.Errorf("failed to find assignment with id %s: %w", assignmentID, err)
	}

	return &assignment, nil
}

func (s *Storage) UpdateAssignment(ctx context.Context, assignmentID string, assignment *model.Assignment) (*model.Assignment, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)
	returnDocumentOption := options.After
	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDocumentOption,
	}
	var updatedAssignment model.Assignment
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": assignmentID}, update(assignment), options).Decode(&updatedAssignment)
	if err != nil {
		return nil, fmt.Errorf("failed to find and update assignment with id %s: %w", assignmentID, err)
	}

	return &updatedAssignment, nil
}

func (s *Storage) DeleteAssignment(ctx context.Context, assignmentID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)
	if _, err := collection.DeleteOne(ctx, bson.M{"_id": assignmentID}); err != nil {
		return fmt.Errorf("failed to delete assignment with id %s: %w", assignmentID, err)
	}

	return nil
}

func update(assignment *model.Assignment) bson.M {
	return bson.M{"$set": assignment}
}
