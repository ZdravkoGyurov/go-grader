package handlers

import (
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type Logout struct {
	Controller *controller.Controller
}

func (h *Logout) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	cookie, err := request.Cookie(h.Controller.Config.SessionCookieName)
	if err != nil {
		response.SendData(writer, http.StatusOK, struct{}{})
		return
	}

	sessionID := cookie.Value
	if err = h.Controller.Logout(ctx, sessionID); err != nil {
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	expiredCookie := http.Cookie{
		Name:    h.Controller.Config.SessionCookieName,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(writer, &expiredCookie)

	log.Info().Printf("logged out user with session id %s\n", sessionID)
	response.SendData(writer, http.StatusOK, struct{}{})
}
