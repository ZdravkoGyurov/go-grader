package handlers

import (
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type Handlers struct {
	Assignment   *AssignmentHandler
	Login        *LoginHandler
	Logout       *LogoutHandler
	Registration *RegistrationHandler
	TestRun      *TestrunHandler
}

func NewHandlers(appContext app.Context, ctrl *controller.Controller) *Handlers {
	return &Handlers{
		Assignment:   &AssignmentHandler{Controller: ctrl},
		Login:        &LoginHandler{Controller: ctrl},
		Logout:       &LogoutHandler{Controller: ctrl},
		Registration: &RegistrationHandler{Controller: ctrl},
		TestRun:      &TestrunHandler{Controller: ctrl},
	}
}
