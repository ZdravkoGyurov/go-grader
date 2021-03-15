package api

import (
	"github.com/ZdravkoGyurov/go-grader/api/handlers/account/login"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/account/logout"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/account/registration"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/assignment"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/testrun"
	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db"
	"github.com/ZdravkoGyurov/go-grader/executor"
)

type Handlers struct {
	Assignment   *assignment.HTTPHandler
	Login        *login.HTTPHandler
	Logout       *logout.HTTPHandler
	Registration *registration.HTTPHandler
	TestRun      *testrun.HTTPHandler
}

func NewHandlers(appCtx app.Context, dbHandler db.Handlers, exec *executor.Executor) *Handlers {
	return &Handlers{
		Assignment:   assignment.NewHTTPHandler(dbHandler.Assignment),
		Login:        login.NewHTTPHandler(appCtx, dbHandler.User, dbHandler.Session),
		Logout:       logout.NewHTTPHandler(appCtx, dbHandler.Session),
		Registration: registration.NewHTTPHandler(appCtx, dbHandler.User),
		TestRun:      testrun.NewHTTPHandler(appCtx, exec),
	}
}
