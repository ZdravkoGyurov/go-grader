package api

import (
	"grader/api/handlers/assignment"
	"grader/api/handlers/login"
	"grader/api/handlers/logout"
	"grader/api/handlers/registration"
	"grader/api/handlers/testrun"
	"grader/app"
	"grader/db"
	"grader/executor"
)

type Handler struct {
	Assignment   *assignment.HTTPHandler
	Login        *login.HTTPHandler
	Logout       *logout.HTTPHandler
	Registration *registration.HTTPHandler
	TestRun      *testrun.HTTPHandler
}

func NewHandler(appCtx app.Context, dbHandler db.Handler, exec *executor.Executor) *Handler {
	return &Handler{
		Assignment:   assignment.NewHTTPHandler(dbHandler.Assignment),
		Login:        login.NewHTTPHandler(appCtx, dbHandler.User, dbHandler.Session),
		Logout:       logout.NewHTTPHandler(appCtx, dbHandler.Session),
		Registration: registration.NewHTTPHandler(dbHandler.User),
		TestRun:      testrun.NewHTTPHandler(exec),
	}
}
