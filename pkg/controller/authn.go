package controller

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

func (c *Controller) GetUserBySessionID(ctx context.Context, sessionID string) (*model.User, error) {
	session, err := c.storage.ReadSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	user, err := c.storage.ReadUserByID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
