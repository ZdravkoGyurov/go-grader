package controller

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/ZdravkoGyurov/go-grader/pkg/random"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Login(ctx context.Context, username, password string) (string, error) {
	user, err := c.Storage.ReadUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	session := model.Session{
		ID:     random.String(20),
		UserID: user.ID,
	}

	err = c.Storage.CreateSession(ctx, &session)
	if err != nil {
		return "", err
	}

	log.Info().Printf("logged in user %s with session id %s\n", user.ID, session.ID)
	return session.ID, nil
}
