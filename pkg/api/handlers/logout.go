package handlers

import (
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
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
		response.Send(writer, http.StatusOK, struct{}{}, nil)
		return
	}

	sessionID := cookie.Value
	if err = h.Controller.Logout(ctx, sessionID); err != nil {
		response.Send(writer, http.StatusInternalServerError, nil, err)
		return
	}

	expiredCookie := http.Cookie{
		Name:    h.Controller.Config.SessionCookieName,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(writer, &expiredCookie)

	log.Info().Printf("logged out user with session id %s\n", sessionID)
	response.Send(writer, http.StatusOK, struct{}{}, nil)
}
