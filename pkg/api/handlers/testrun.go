package handlers

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type Submission struct {
	Controller *controller.Controller
}

func (h *Submission) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	_, err := h.Controller.CreateSubmission()
	if err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	response.SendData(writer, http.StatusOK, struct{}{})
}
