package api

import (
	"context"
	"encoding/json"
	"grader/db/models"
	"net/http"
)

type assignmentsDBHandler interface {
	CreateAssignment(ctx context.Context, assignment models.Assignment) error
}

// AssignmentsHandler ...
type AssignmentsHandler struct {
	dbHandler assignmentsDBHandler
}

// NewAssignmentsHandler creates a new assignments http handler
func NewAssignmentsHandler(dbHandler assignmentsDBHandler) *AssignmentsHandler {
	return &AssignmentsHandler{
		dbHandler: dbHandler,
	}
}

// Post ...
func (h *AssignmentsHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var assignment models.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&assignment); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	if err := h.dbHandler.CreateAssignment(ctx, assignment); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
