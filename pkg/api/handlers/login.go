package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type Login struct {
	Controller *controller.Controller
}

func (h *Login) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.Controller.Config.SessionCookieName); err == nil {
		response.SendData(writer, http.StatusOK, struct{}{})
		return
	}

	username, password, ok := request.BasicAuth()
	if !ok {
		err := errors.New("failed to get username and password from authorization header")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	sessionID, err := h.Controller.Login(ctx, username, password)
	if err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	// TODO: make secure cookie
	cookie := &http.Cookie{
		Name:     h.Controller.Config.SessionCookieName,
		Value:    sessionID,
		Domain:   "localhost",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(writer, cookie)

	response.SendData(writer, http.StatusOK, struct{}{})
}
