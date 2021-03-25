package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

type Registration struct {
	Controller *controller.Controller
}

func (h *Registration) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.Controller.Config.SessionCookieName); err == nil {
		err := errors.New("failed to register logged in user")
		response.SendError(writer, http.StatusUnprocessableEntity, err)
		return
	}

	username, password, ok := request.BasicAuth()
	if !ok {
		err := errors.New("failed to get username and password from authorization header")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	var user model.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		err = errors.Wrap(err, "failed to decode user from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}
	user.Username = username
	user.Password = password

	if err := h.Controller.Register(ctx, &user); err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	response.SendData(writer, http.StatusOK, struct{}{})
}
