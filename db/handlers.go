package db

import (
	"grader/app"
	"grader/db/handlers/assignment"
	"grader/db/handlers/session"
	"grader/db/handlers/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handlers struct {
	Assignment *assignment.DBHandler
	Session    *session.DBHandler
	User       *user.DBHandler
}

func NewHandlers(appCtx app.Context, client *mongo.Client) *Handlers {
	return &Handlers{
		Assignment: assignment.NewDBHandler(appCtx, client),
		Session:    session.NewDBHandler(appCtx, client),
		User:       user.NewDBHandler(appCtx, client),
	}
}
