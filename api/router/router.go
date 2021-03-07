package router

import (
	"net/http"

	"grader/api"

	"github.com/gorilla/mux"
)

// New creates a mux router with configured routes
func New(assignmentsHTTPHandler *api.AssignmentsHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/assignments", assignmentsHTTPHandler.Post).Methods(http.MethodPost)

	return r
}
