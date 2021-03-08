package router

import (
	"net/http"

	"grader/api"
	"grader/api/router/paths"

	"github.com/gorilla/mux"
)

// New creates a mux router with configured routes
func New(assignmentsHTTPHandler *api.AssignmentsHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(paths.Assignments, assignmentsHTTPHandler.Post).Methods(http.MethodPost)
	r.HandleFunc(paths.AssignmentsWithID, assignmentsHTTPHandler.Get).Methods(http.MethodGet)
	r.HandleFunc(paths.AssignmentsWithID, assignmentsHTTPHandler.Patch).Methods(http.MethodPatch)
	r.HandleFunc(paths.AssignmentsWithID, assignmentsHTTPHandler.Delete).Methods(http.MethodDelete)

	return r
}
