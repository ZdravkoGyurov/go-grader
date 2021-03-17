package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type logoutStorage interface {
	DeleteSession(ctx context.Context, sessionID string) error
}

type logoutHandler struct {
	appContext app.Context
	logoutStorage
}

func (h *logoutHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.appContext.Cfg.SessionCookieName); err != nil {
		log.Warning().Println(errors.New("failed to logout logged out user"))
		writer.WriteHeader(http.StatusOK)
		return
	}

	cookie, err := request.Cookie("Grader")
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionID := cookie.Value
	if err := h.logoutStorage.DeleteSession(ctx, sessionID); err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	expiredCookie := http.Cookie{
		Name:    h.appContext.Cfg.SessionCookieName,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(writer, &expiredCookie)

	log.Info().Printf("logged out user with session id %s\n", sessionID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
