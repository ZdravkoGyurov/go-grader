package db

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZdravkoGyurov/go-grader/internal/app"
	"github.com/ZdravkoGyurov/go-grader/internal/db/handlers/assignment"
	"github.com/ZdravkoGyurov/go-grader/internal/db/handlers/session"
	"github.com/ZdravkoGyurov/go-grader/internal/db/handlers/user"
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
