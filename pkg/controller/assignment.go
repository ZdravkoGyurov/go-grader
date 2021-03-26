package controller

import (
	"context"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/google/uuid"
)

func (c *Controller) CreateAssignment(ctx context.Context, assignment *model.Assignment) error {
	assignment.ID = uuid.NewString()
	assignment.CreatedOn = time.Now()
	assignment.LastUpdatedOn = time.Now()

	if err := c.storage.CreateAssignment(ctx, assignment); err != nil {
		return err
	}

	return nil
}

func (c *Controller) GetAllAssignments(ctx context.Context, courseID string) ([]*model.Assignment, error) {
	if courseID == "" {
		return nil, errors.Wrap(errors.ErrInvalidInput, "course id cannot be empty")
	}

	assignments, err := c.storage.ReadAllAssignments(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}

func (c *Controller) GetAssignment(ctx context.Context, assignmentID string) (*model.Assignment, error) {
	assignment, err := c.storage.ReadAssignment(ctx, assignmentID)
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (c *Controller) UpdateAssignment(ctx context.Context, assignmentID string, assignment *model.Assignment) (*model.Assignment, error) {
	assignment.LastUpdatedOn = time.Now()

	updatedAssignment, err := c.storage.UpdateAssignment(ctx, assignmentID, assignment)
	if err != nil {
		return nil, err
	}

	return updatedAssignment, nil
}

func (c *Controller) DeleteAssignment(ctx context.Context, assignmentID string) error {
	if err := c.storage.DeleteAssignment(ctx, assignmentID); err != nil {
		return err
	}

	return nil
}
