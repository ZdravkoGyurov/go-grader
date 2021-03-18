package controller

import "context"

func (c *Controller) Logout(ctx context.Context, sessionID string) error {
	if err := c.storage.DeleteSession(ctx, sessionID); err != nil {
		return err
	}

	return nil
}
