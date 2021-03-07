package db

import (
	"context"
	"fmt"
	"grader/db/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// AssignmentsHandler ...
type AssignmentsHandler struct {
	client *mongo.Client
}

// NewAssignmentsHandler creates a new assignments DB handler
func NewAssignmentsHandler(client *mongo.Client) *AssignmentsHandler {
	return &AssignmentsHandler{
		client: client,
	}
}

// CreateAssignment creates new assignment in the database
func (h *AssignmentsHandler) CreateAssignment(ctx context.Context, assignment models.Assignment) error {
	collection := h.client.Database("grader").Collection(models.AssignmentsCollectionName)
	if _, err := collection.InsertOne(ctx, assignment); err != nil {
		return fmt.Errorf("failed to insert assignment: %w", err)
	}

	return nil
}
