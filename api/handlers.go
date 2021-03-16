package api

import (
	"github.com/ZdravkoGyurov/go-grader/api/handlers/account/login"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/account/logout"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/account/registration"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/assignment"
	"github.com/ZdravkoGyurov/go-grader/api/handlers/testrun"
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

type Handlers struct {
	Assignment   *assignment.HTTPHandler
	Login        *login.HTTPHandler
	Logout       *logout.HTTPHandler
	Registration *registration.HTTPHandler
	TestRun      *testrun.HTTPHandler
}

func NewHandlers(appContext app.Context, storage *storage.Storage, exec *executor.Executor) *Handlers {
	return &Handlers{
		Assignment:   assignment.NewHTTPHandler(storage),
		Login:        login.NewHTTPHandler(appContext, storage),
		Logout:       logout.NewHTTPHandler(appContext, storage),
		Registration: registration.NewHTTPHandler(appContext, storage),
		TestRun:      testrun.NewHTTPHandler(appContext, exec),
	}
}
