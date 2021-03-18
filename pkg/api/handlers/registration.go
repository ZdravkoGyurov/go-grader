package handlers

import (
	"errors"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type RegistrationHandler struct {
	Controller *controller.Controller
}

func (h *RegistrationHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.Controller.Config.SessionCookieName); err == nil {
		log.Error().Println(errors.New("failed to register logged in user"))
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	username, password, ok := request.BasicAuth()
	if !ok {
		log.Error().Println("failed to get username and password from authorization header")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Controller.Register(ctx, username, password); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
