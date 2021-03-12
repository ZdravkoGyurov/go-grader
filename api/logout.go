package api

import (
	"context"
	"grader/app"
	"grader/log"
	"net/http"
	"time"
)

type sessionHandler interface {
	DeleteSession(ctx context.Context, sessionID string) error
}

// LogoutHandler ...
type LogoutHandler struct {
	appCtx app.Context
	sessionHandler
}

// NewLogoutHandler creates a new logout http handler
func NewLogoutHandler(appCtx app.Context, sessionHandler sessionHandler) *LogoutHandler {
	return &LogoutHandler{
		appCtx:         appCtx,
		sessionHandler: sessionHandler,
	}
}

// Post ...
func (h *LogoutHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	cookie, err := request.Cookie("Grader")
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionID := cookie.Value
	if err := h.sessionHandler.DeleteSession(ctx, sessionID); err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	expiredCookie := http.Cookie{
		Name:    h.appCtx.Cfg.SessionCookieName,
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(writer, &expiredCookie)

	log.Info().Printf("logged out user with session id %s\n", sessionID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
