package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

type AssignmentHandler struct {
	Controller *controller.Controller
}

func (h *AssignmentHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var assignment model.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&assignment); err != nil {
		err = fmt.Errorf("failed to decode assignment from request body: %s", err)
		response.Send(writer, http.StatusBadRequest, nil, err)
		return
	}

	if err := h.Controller.CreateAssignment(ctx, &assignment); err != nil {
		response.Send(writer, http.StatusInternalServerError, nil, err)
		return
	}

	log.Info().Printf("created assignment with id %s\n", assignment.ID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (h *AssignmentHandler) Get(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		err := errors.New("failed to get assignment id path parameter")
		response.Send(writer, http.StatusInternalServerError, nil, err)
		return
	}

	assignment, err := h.Controller.GetAssignment(ctx, assignmentID)
	if err != nil {
		response.Send(writer, http.StatusInternalServerError, nil, err)
		return
	}

	response.Send(writer, http.StatusOK, assignment, nil)
}

func (h *AssignmentHandler) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		log.Error().Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updateAssignment model.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updateAssignment); err != nil {
		log.Error().Printf("failed to decode assignment from request body: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updateAssignment.ID = ""

	updatedAssignment, err := h.Controller.UpdateAssignment(ctx, assignmentID, &updateAssignment)
	if err != nil {
		log.Error().Printf("failed to marshal assignment json data: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(updatedAssignment)
	if err != nil {
		log.Error().Printf("failed to marshal assignment json data: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Printf("updated assignment with id %s\n", assignmentID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}

func (h *AssignmentHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		log.Error().Println("failed to get assignment id path parameter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.Controller.DeleteAssignment(ctx, assignmentID); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError) // handle error method
		return
	}

	log.Info().Printf("deleted assignment with id %s\n", assignmentID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}
