package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/req"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/gorilla/mux"
)

type Submission struct {
	Controller *controller.Controller
}

func (h *Submission) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var submission model.Submission
	if err := json.NewDecoder(request.Body).Decode(&submission); err != nil {
		err = errors.Wrap(err, "failed to decode submission from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	reqData, ok := req.GetRequestData(request)
	if !ok {
		err := errors.New("failed to retrieve github username")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	if err := h.Controller.CreateSubmission(ctx, &submission, reqData.GithubUsername); err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, submission)
}

func (h *Submission) GetAll(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	body := struct {
		UserID       string `json:"user_id"`
		AssignmentID string `json:"assignment_id"`
	}{}
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		err = errors.Wrap(err, "failed to decode user_id or assignment_id from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	submissions, err := h.Controller.GetAllSubmissions(ctx, body.UserID, body.AssignmentID)
	if err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, submissions)
}

func (h *Submission) Get(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	submissionID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get submission id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	submission, err := h.Controller.GetSubmission(ctx, submissionID)
	if err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, submission)
}

func (h *Submission) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	submissionID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get submission id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	var updateSubmission model.Submission
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updateSubmission); err != nil {
		err = errors.Wrap(err, "failed to decode submission from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}
	updateSubmission.ID = ""

	updatedSubmission, err := h.Controller.UpdateSubmission(ctx, submissionID, &updateSubmission)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal submission json data")
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	log.Info().Printf("updated submission with id %s\n", submissionID)
	response.SendData(writer, http.StatusOK, updatedSubmission)
}

func (h *Submission) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	submissionID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get submission id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	if err := h.Controller.DeleteSubmission(ctx, submissionID); err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	log.Info().Printf("deleted submission with id %s\n", submissionID)
	response.SendData(writer, http.StatusNoContent, struct{}{})
}
