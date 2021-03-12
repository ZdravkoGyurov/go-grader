package router

import (
	"net/http"

	"grader/api"
	"grader/api/router/paths"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	Registration *api.RegistrationHandler
	Login        *api.LoginHandler
	Logout       *api.LogoutHandler
	Assignments  *api.AssignmentsHandler
	TestRun      *api.TestRunHandler
}

// New creates a mux router with configured routes
func New(httpHandlers HTTPHandlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(paths.Register, httpHandlers.Registration.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Login, httpHandlers.Login.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Logout, httpHandlers.Logout.Post).Methods(http.MethodPost)

	r.HandleFunc(paths.Assignments, httpHandlers.Assignments.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.AssignmentsWithID, httpHandlers.Assignments.Get).Methods(http.MethodGet)
	r.HandleFunc(paths.AssignmentsWithID, httpHandlers.Assignments.Patch).Methods(http.MethodPatch)
	r.HandleFunc(paths.AssignmentsWithID, httpHandlers.Assignments.Delete).Methods(http.MethodDelete)

	r.HandleFunc(paths.TestRun, httpHandlers.TestRun.Post).Methods(http.MethodPost)

	return r
}
