package controller

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/ZdravkoGyurov/go-grader/pkg/api"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

func (c *Controller) Register(ctx context.Context, userBody *model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user := model.User{
		ID:             uuid.NewString(),
		Username:       userBody.Username,
		Fullname:       userBody.Fullname,
		GithubUsername: userBody.GithubUsername,
		Password:       string(hash),
		Permissions:    api.StudentPermissions,
		CourseIDs:      []string{},
	}

	if err := c.storage.CreateUser(ctx, &user); err != nil {
		return err
	}

	log.Info().Printf("created user with id %s\n", user.ID)
	return nil
}
