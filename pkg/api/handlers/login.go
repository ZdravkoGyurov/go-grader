package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type LoginHandler struct {
	Controller *controller.Controller
}

func (h *LoginHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.Controller.Config.SessionCookieName); err == nil {
		log.Warning().Println(errors.New("failed to login logged in user"))
		writer.WriteHeader(http.StatusOK)
		return
	}

	username, password, ok := request.BasicAuth()
	if !ok {
		log.Error().Println("failed to get username and password from authorization header")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionID, err := h.Controller.Login(ctx, username, password)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError) // handle error method
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
	fmt.Printf("created cookie: %+v\n", *cookie)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
