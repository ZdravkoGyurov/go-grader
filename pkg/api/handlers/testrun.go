package handlers

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type TestrunHandler struct {
	Controller *controller.Controller
}

func (h *TestrunHandler) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	_, err := h.Controller.CreateTestrun()
	if err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	response.SendData(writer, http.StatusOK, struct{}{})
}
