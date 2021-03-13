package db

import (
	"grader/app"
	"grader/db/handlers/assignment"
	"grader/db/handlers/session"
	"grader/db/handlers/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Assignment *assignment.DBHandler
	Session    *session.DBHandler
	User       *user.DBHandler
}

func NewHandler(appCtx app.Context, client *mongo.Client) *Handler {
	return &Handler{
		Assignment: assignment.NewDBHandler(appCtx, client),
		Session:    session.NewDBHandler(appCtx, client),
		User:       user.NewDBHandler(appCtx, client),
	}
}
