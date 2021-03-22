package controller

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

func (c *Controller) GetAllUsers(ctx context.Context, courseID string) ([]*model.User, error) {
	if courseID == "" {
		return nil, errors.Wrap(errors.ErrInvalidInput, "course id cannot be empty")
	}

	users, err := c.storage.ReadAllUsers(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Controller) UpdateUser(ctx context.Context, userID string, user *model.User) (*model.User, error) {
	if user.Password != "" {
		return nil, errors.Wrap(errors.ErrInvalidInput, "cannot change user password")
	}

	updatedUser, err := c.storage.UpdateUser(ctx, userID, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (c *Controller) DeleteUser(ctx context.Context, userID string) error {
	if err := c.storage.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}
