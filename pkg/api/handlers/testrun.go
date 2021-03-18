package handlers

import (
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type TestrunHandler struct {
	Controller *controller.Controller
}

func (h *TestrunHandler) Post(writer http.ResponseWriter, request *http.Request) {
	// ctx := request.Context()

	jobID, err := h.Controller.CreateTestrun()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"jobId": %s}`, jobID)))
}
