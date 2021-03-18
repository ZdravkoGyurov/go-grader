package controller

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/google/uuid"
)

func (c *Controller) CreateAssignment(ctx context.Context, assignment *model.Assignment) error {
	assignment.ID = uuid.NewString()

	if err := c.Storage.CreateAssignment(ctx, assignment); err != nil {
		return err
	}

	return nil
}

func (c *Controller) GetAssignment(ctx context.Context, assignmentID string) (*model.Assignment, error) {
	assignment, err := c.Storage.ReadAssignment(ctx, assignmentID)
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (c *Controller) UpdateAssignment(ctx context.Context, assignmentID string, assignment *model.Assignment) (*model.Assignment, error) {
	updatedAssignment, err := c.Storage.UpdateAssignment(ctx, assignmentID, assignment)
	if err != nil {
		return nil, err
	}

	return updatedAssignment, nil
}

func (c *Controller) DeleteAssignment(ctx context.Context, assignmentID string) error {
	if err := c.Storage.DeleteAssignment(ctx, assignmentID); err != nil {
		return err
	}

	return nil
}
