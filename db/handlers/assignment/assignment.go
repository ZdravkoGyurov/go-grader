package assignment

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db/models"
)

// DBHandler ...
type DBHandler struct {
	collection *mongo.Collection
}

// NewDBHandler creates a new assignment DB handler
func NewDBHandler(appCtx app.Context, client *mongo.Client) *DBHandler {
	return &DBHandler{
		collection: client.Database(appCtx.Cfg.DatabaseName).
			Collection(models.AssignmentCollectionName),
	}
}

// Create creates new assignment in the database
func (h *DBHandler) Create(ctx context.Context, assignment *models.Assignment) error {
	_, err := h.collection.InsertOne(ctx, assignment)
	if err != nil {
		return fmt.Errorf("failed to insert assignment: %w", err)
	}

	return nil
}

// Read ...
func (h *DBHandler) Read(ctx context.Context, assignmentID string) (*models.Assignment, error) {
	var assignment models.Assignment
	if err := h.collection.FindOne(ctx, bson.M{"_id": assignmentID}).Decode(&assignment); err != nil {
		return nil, fmt.Errorf("failed to find assignment with id %s: %w", assignmentID, err)
	}

	return &assignment, nil
}

// Update ...
func (h *DBHandler) Update(ctx context.Context, assignmentID string, assignment *models.Assignment) (*models.Assignment, error) {
	returnDocumentOption := options.After
	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDocumentOption,
	}
	var updatedAssignment models.Assignment
	err := h.collection.FindOneAndUpdate(ctx, bson.M{"_id": assignmentID}, update(assignment), options).Decode(&updatedAssignment)
	if err != nil {
		return nil, fmt.Errorf("failed to find and update assignment with id %s: %w", assignmentID, err)
	}

	return &updatedAssignment, nil
}

// Delete ...
func (h *DBHandler) Delete(ctx context.Context, assignmentID string) error {
	if _, err := h.collection.DeleteOne(ctx, bson.M{"_id": assignmentID}); err != nil {
		return fmt.Errorf("failed to delete assignment with id %s: %w", assignmentID, err)
	}

	return nil
}

func update(assignment *models.Assignment) bson.M {
	return bson.M{"$set": assignment}
}
