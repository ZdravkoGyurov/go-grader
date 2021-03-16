package assignment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

type assignmentStorage interface {
	CreateAssignment(ctx context.Context, assignment *model.Assignment) error
	ReadAssignment(ctx context.Context, assignmentID string) (*model.Assignment, error)
	UpdateAssignment(ctx context.Context, assignmentID string, assignment *model.Assignment) (*model.Assignment, error)
	DeleteAssignment(ctx context.Context, assignmentID string) error
}

// HTTPHandler ...
type HTTPHandler struct {
	assignmentStorage assignmentStorage
}

// NewHTTPHandler creates a new assignment http handler
func NewHTTPHandler(assignmentStorage assignmentStorage) *HTTPHandler {
	return &HTTPHandler{
		assignmentStorage: assignmentStorage,
	}
}

// Post ...
func (h *HTTPHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var assignment model.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&assignment); err != nil {
		log.Info().Printf("failed to decode assignment from request body: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	assignment.ID = uuid.NewString()

	if err := h.assignmentStorage.CreateAssignment(ctx, &assignment); err != nil {
		log.Info().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Printf("created assignment with id %s\n", assignment.ID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

// Get ...
func (h *HTTPHandler) Get(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		log.Info().Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	assignment, err := h.assignmentStorage.ReadAssignment(ctx, assignmentID)
	if err != nil {
		log.Info().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(assignment)
	if err != nil {
		log.Info().Printf("failed to marshal assignment json data: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}

// Patch ...
func (h *HTTPHandler) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		log.Info().Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateAssignment model.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updateAssignment); err != nil {
		log.Info().Printf("failed to decode assignment from request body: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updateAssignment.ID = ""

	updatedAssignment, err := h.assignmentStorage.UpdateAssignment(ctx, assignmentID, &updateAssignment)
	if err != nil {
		log.Info().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(updatedAssignment)
	if err != nil {
		log.Info().Printf("failed to marshal assignment json data: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Printf("updated assignment with id %s\n", assignmentID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}

// Delete ...
func (h *HTTPHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		log.Info().Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.assignmentStorage.DeleteAssignment(ctx, assignmentID); err != nil {
		log.Info().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Printf("deleted assignment with id %s\n", assignmentID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}
