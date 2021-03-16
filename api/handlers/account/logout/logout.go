package logout

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ZdravkoGyurov/go-grader/api/handlers/account"
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

type logoutStorage interface {
	DeleteSession(ctx context.Context, sessionID string) error
}

// HTTPHandler ...
type HTTPHandler struct {
	appContext app.Context
	logoutStorage
}

// NewHTTPHandler creates a new logout http handler
func NewHTTPHandler(appContext app.Context, logoutStorage logoutStorage) *HTTPHandler {
	return &HTTPHandler{
		appContext:    appContext,
		logoutStorage: logoutStorage,
	}
}

// Post ...
func (h *HTTPHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if !account.UserLoggedIn(h.appContext, request) {
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
