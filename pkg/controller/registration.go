package controller

import (
	"context"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Register(ctx context.Context, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user := model.User{
		ID:          uuid.NewString(),
		Username:    username,
		Fullname:    "fullname", // TODO: get from body
		Password:    string(hash),
		Permissions: []string{"STUDENT"}, // TODO: add permissions
		Disabled:    false,
	}

	if err := c.Storage.CreateUser(ctx, &user); err != nil {
		return err
	}

	log.Info().Printf("created user with id %s\n", user.ID)
	return nil
}
