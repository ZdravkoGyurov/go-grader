package router

import (
	"net/http"

	"grader/api"
	"grader/api/router/paths"

	"github.com/gorilla/mux"
)

// New creates a mux router with configured routes
func New(httpHandler *api.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(paths.Register, httpHandler.Registration.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Login, httpHandler.Login.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.Logout, httpHandler.Logout.Post).Methods(http.MethodPost)

	r.HandleFunc(paths.Assignments, httpHandler.Assignment.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.AssignmentsWithID, httpHandler.Assignment.Get).Methods(http.MethodGet)
	r.HandleFunc(paths.AssignmentsWithID, httpHandler.Assignment.Patch).Methods(http.MethodPatch)
	r.HandleFunc(paths.AssignmentsWithID, httpHandler.Assignment.Delete).Methods(http.MethodDelete)

	r.HandleFunc(paths.TestRun, httpHandler.TestRun.Post).Methods(http.MethodPost)

	return r
}
