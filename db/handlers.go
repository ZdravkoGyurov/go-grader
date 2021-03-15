package db

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db/handlers/assignment"
	"github.com/ZdravkoGyurov/go-grader/db/handlers/session"
	"github.com/ZdravkoGyurov/go-grader/db/handlers/user"
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
