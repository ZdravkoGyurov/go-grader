package logout

import (
	"context"
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/log"
)

type sessionDBHandler interface {
	Delete(ctx context.Context, sessionID string) error
}

// HTTPHandler ...
type HTTPHandler struct {
	appCtx app.Context
	sessionDBHandler
}

// NewHTTPHandler creates a new logout http handler
func NewHTTPHandler(appCtx app.Context, sessionDBHandler sessionDBHandler) *HTTPHandler {
	return &HTTPHandler{
		appCtx:           appCtx,
		sessionDBHandler: sessionDBHandler,
	}
}

// Post ...
func (h *HTTPHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	cookie, err := request.Cookie("Grader")
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionID := cookie.Value
	if err := h.sessionDBHandler.Delete(ctx, sessionID); err != nil {
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
