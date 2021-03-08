package db

import (
	"context"
	"fmt"
	"grader/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AssignmentsHandler ...
type AssignmentsHandler struct {
	collection *mongo.Collection
}

// NewAssignmentsHandler creates a new assignments DB handler
func NewAssignmentsHandler(client *mongo.Client) *AssignmentsHandler {
	return &AssignmentsHandler{
		collection: client.Database("grader").Collection(models.AssignmentsCollectionName),
	}
}

// CreateAssignment creates new assignment in the database
func (h *AssignmentsHandler) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	result, err := h.collection.InsertOne(ctx, assignment)
	if err != nil {
		return fmt.Errorf("failed to insert assignment: %w", err)
	}

	assignment.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

// ReadAssignment ...
func (h *AssignmentsHandler) ReadAssignment(ctx context.Context, assignmentID string) (*models.Assignment, error) {
	filter, err := filterByID(assignmentID)
	if err != nil {
		return nil, err
	}
	var assignment models.Assignment
	if err := h.collection.FindOne(ctx, filter).Decode(&assignment); err != nil {
		return nil, fmt.Errorf("failed to find assignment with id %s: %w", assignmentID, err)
	}

	return &assignment, nil
}

// UpdateAssignment ...
func (h *AssignmentsHandler) UpdateAssignment(ctx context.Context, assignmentID string, assignment *models.Assignment) (*models.Assignment, error) {
	filter, err := filterByID(assignmentID)
	if err != nil {
		return nil, err
	}
	returnDocumentOption := options.After
	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDocumentOption,
	}
	var updatedAssignment models.Assignment
	err = h.collection.FindOneAndUpdate(ctx, filter, update(assignment), options).Decode(&updatedAssignment)
	if err != nil {
		return nil, fmt.Errorf("failed to find and update assignment with id %s: %w", assignmentID, err)
	}

	return &updatedAssignment, nil
}

// DeleteAssignment ...
func (h *AssignmentsHandler) DeleteAssignment(ctx context.Context, assignmentID string) error {
	filter, err := filterByID(assignmentID)
	if err != nil {
		return err
	}
	if _, err := h.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("failed to delete assignment with id %s: %w", assignmentID, err)
	}

	return nil
}

func filterByID(assignmentID string) (bson.M, error) {
	objectID, err := primitive.ObjectIDFromHex(assignmentID)
	if err != nil {
		return bson.M{}, fmt.Errorf("failed to generate filter for id %s: %w", assignmentID, err)
	}

	return bson.M{"_id": objectID}, nil
}

func update(assignment *models.Assignment) bson.M {
	return bson.M{"$set": assignment}
}
