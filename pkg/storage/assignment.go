package storage

import (
	"context"
	"fmt"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
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
	if err := collection.FindOne(ctx, filterByID(assignmentID)).Decode(&assignment); err != nil {
		return nil, fmt.Errorf("failed to find assignment with id %s: %w", assignmentID, err)
	}

	return &assignment, nil
}

func (s *Storage) ReadAllAssignments(ctx context.Context, courseID string) ([]*model.Assignment, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)

	cursor, err := collection.Find(ctx, filterAssignmentByCourseID(courseID))
	if err != nil {
		return nil, fmt.Errorf("failed to find all assignments with course_id %s: %w", courseID, err)
	}

	var assignments []*model.Assignment
	if err = cursor.All(ctx, &assignments); err != nil {
		return nil, fmt.Errorf("failed to decode all assignments with course_id %s: %w", courseID, err)
	}

	return assignments, nil
}

func (s *Storage) UpdateAssignment(ctx context.Context, assignmentID string, assignment *model.Assignment) (*model.Assignment, error) {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)
	var updatedAssignment model.Assignment
	result := collection.FindOneAndUpdate(ctx, filterByID(assignmentID), update(assignment), updateOpts())
	if err := result.Decode(&updatedAssignment); err != nil {
		return nil, fmt.Errorf("failed to find and update assignment with id %s: %w", assignmentID, err)
	}

	return &updatedAssignment, nil
}

func (s *Storage) DeleteAssignment(ctx context.Context, assignmentID string) error {
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection(assignmentCollection)
	if _, err := collection.DeleteOne(ctx, filterByID(assignmentID)); err != nil {
		return fmt.Errorf("failed to delete assignment with id %s: %w", assignmentID, err)
	}
	return nil
}
