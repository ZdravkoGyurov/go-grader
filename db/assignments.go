package db

import (
	"context"
	"fmt"
	"grader/app"
	"grader/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AssignmentsHandler ...
type AssignmentsHandler struct {
	collection *mongo.Collection
}

// NewAssignmentsHandler creates a new assignments DB handler
func NewAssignmentsHandler(appCtx app.Context, client *mongo.Client) *AssignmentsHandler {
	return &AssignmentsHandler{
		collection: client.Database(appCtx.Cfg.DatabaseName).
			Collection(models.AssignmentsCollectionName),
	}
}

// CreateAssignment creates new assignment in the database
func (h *AssignmentsHandler) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	_, err := h.collection.InsertOne(ctx, assignment)
	if err != nil {
		return fmt.Errorf("failed to insert assignment: %w", err)
	}

	return nil
}

// ReadAssignment ...
func (h *AssignmentsHandler) ReadAssignment(ctx context.Context, assignmentID string) (*models.Assignment, error) {
	var assignment models.Assignment
	if err := h.collection.FindOne(ctx, filterByID(assignmentID)).Decode(&assignment); err != nil {
		return nil, fmt.Errorf("failed to find assignment with id %s: %w", assignmentID, err)
	}

	return &assignment, nil
}

// UpdateAssignment ...
func (h *AssignmentsHandler) UpdateAssignment(ctx context.Context, assignmentID string, assignment *models.Assignment) (*models.Assignment, error) {
	returnDocumentOption := options.After
	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDocumentOption,
	}
	var updatedAssignment models.Assignment
	err := h.collection.FindOneAndUpdate(ctx, filterByID(assignmentID), update(assignment), options).Decode(&updatedAssignment)
	if err != nil {
		return nil, fmt.Errorf("failed to find and update assignment with id %s: %w", assignmentID, err)
	}

	return &updatedAssignment, nil
}

// DeleteAssignment ...
func (h *AssignmentsHandler) DeleteAssignment(ctx context.Context, assignmentID string) error {
	if _, err := h.collection.DeleteOne(ctx, filterByID(assignmentID)); err != nil {
		return fmt.Errorf("failed to delete assignment with id %s: %w", assignmentID, err)
	}

	return nil
}

func filterByID(id string) bson.M {
	return bson.M{"_id": id}
}

func update(assignment *models.Assignment) bson.M {
	return bson.M{"$set": assignment}
}
