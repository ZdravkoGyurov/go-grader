package handlers

import (
	"errors"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type RegistrationHandler struct {
	Controller *controller.Controller
}

func (h *RegistrationHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.Controller.Config.SessionCookieName); err == nil {
		err := errors.New("failed to register logged in user")
		response.Send(writer, http.StatusUnprocessableEntity, nil, err)
		return
	}

	username, password, ok := request.BasicAuth()
	if !ok {
		err := errors.New("failed to get username and password from authorization header")
		response.Send(writer, http.StatusBadRequest, nil, err)
		return
	}

	if err := h.Controller.Register(ctx, username, password); err != nil {
		response.Send(writer, http.StatusInternalServerError, nil, err)
		return
	}

	response.Send(writer, http.StatusOK, struct{}{}, nil)
}
