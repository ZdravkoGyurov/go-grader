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

type Assignment struct {
	Controller *controller.Controller
}

func (h *Assignment) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var assignment model.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&assignment); err != nil {
		err = fmt.Errorf("failed to decode assignment from request body: %s", err)
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	if err := h.Controller.CreateAssignment(ctx, &assignment); err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	log.Info().Printf("created assignment with id %s\n", assignment.ID)
	response.SendData(writer, http.StatusOK, assignment)
}

func (h *Assignment) Get(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		err := errors.New("failed to get assignment id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	assignment, err := h.Controller.GetAssignment(ctx, assignmentID)
	if err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	response.SendData(writer, http.StatusOK, assignment)
}

func (h *Assignment) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		err := errors.New("failed to get assignment id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	var updateAssignment model.Assignment
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updateAssignment); err != nil {
		err = fmt.Errorf("failed to decode assignment from request body: %s", err)
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}
	updateAssignment.ID = ""

	updatedAssignment, err := h.Controller.UpdateAssignment(ctx, assignmentID, &updateAssignment)
	if err != nil {
		err = fmt.Errorf("failed to marshal assignment json data: %s", err)
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	log.Info().Printf("updated assignment with id %s\n", assignmentID)
	response.SendData(writer, http.StatusOK, updatedAssignment)
}

func (h *Assignment) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	assignmentID, ok := mux.Vars(request)[paths.AssignmentIDParam]
	if !ok {
		err := errors.New("failed to get assignment id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	if err := h.Controller.DeleteAssignment(ctx, assignmentID); err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	log.Info().Printf("deleted assignment with id %s\n", assignmentID)
	response.SendData(writer, http.StatusNoContent, struct{}{})
}
