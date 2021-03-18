package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type LogoutHandler struct {
	Controller *controller.Controller
}

func (h *LogoutHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	cookie, err := request.Cookie(h.Controller.Config.SessionCookieName)
	if err != nil {
		log.Warning().Println(errors.New("failed to logout logged out user"))
		writer.WriteHeader(http.StatusOK)
		return
	}

	sessionID := cookie.Value
	if err = h.Controller.Logout(ctx, sessionID); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError) // handle error method
		return
	}

	expiredCookie := http.Cookie{
		Name:    h.Controller.Config.SessionCookieName,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(writer, &expiredCookie)

	log.Info().Printf("logged out user with session id %s\n", sessionID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
