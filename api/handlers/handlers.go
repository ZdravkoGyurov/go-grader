package handlers

import (
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/executor"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

type Handlers struct {
	Assignment   *assignmentHandler
	Login        *loginHandler
	Logout       *logoutHandler
	Registration *registrationHandler
	TestRun      *testrunHandler
}

func NewHandlers(appContext app.Context, storage *storage.Storage, exe *executor.Executor) *Handlers {
	return &Handlers{
		Assignment:   &assignmentHandler{assignmentStorage: storage},
		Login:        &loginHandler{appContext: appContext, loginStorage: storage},
		Logout:       &logoutHandler{appContext: appContext, logoutStorage: storage},
		Registration: &registrationHandler{appContext: appContext, registrationStorage: storage},
		TestRun:      &testrunHandler{appContext: appContext, exe: exe},
	}
}
