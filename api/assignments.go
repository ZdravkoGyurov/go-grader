package api

import (
	"context"
	"encoding/json"
	"grader/api/router/paths"
	"grader/db/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type assignmentsDBHandler interface {
	CreateAssignment(ctx context.Context, assignment *models.Assignment) error
	ReadAssignment(ctx context.Context, assignmentID string) (*models.Assignment, error)
	UpdateAssignment(ctx context.Context, assignmentID string, assignment *models.Assignment) (*models.Assignment, error)
	DeleteAssignment(ctx context.Context, assignmentID string) error
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
		log.Printf("failed to decode assignment from request body: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	assignment.ID = "" // force mongo to generate ID

	if err := h.dbHandler.CreateAssignment(ctx, &assignment); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("created assignment with id %s\n", assignment.ID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

// Get ...
func (h *AssignmentsHandler) Get(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentsIDParam]
	if !ok {
		log.Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	assignment, err := h.dbHandler.ReadAssignment(ctx, assignmentID)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(assignment)
	if err != nil {
		log.Printf("failed to marshal assignment json data: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}

// Patch ...
func (h *AssignmentsHandler) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentsIDParam]
	if !ok {
		log.Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateAssignment models.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updateAssignment); err != nil {
		log.Printf("failed to decode assignment from request body: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updateAssignment.ID = ""

	updatedAssignment, err := h.dbHandler.UpdateAssignment(ctx, assignmentID, &updateAssignment)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(updatedAssignment)
	if err != nil {
		log.Printf("failed to marshal assignment json data: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("updated assignment with id %s\n", assignmentID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}

// Delete ...
func (h *AssignmentsHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentsIDParam]
	if !ok {
		log.Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.dbHandler.DeleteAssignment(ctx, assignmentID); err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("deleted assignment with id %s\n", assignmentID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}
