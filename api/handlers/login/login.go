package login

import (
	"context"
	"grader/app"
	"grader/db/models"
	"grader/log"
	"grader/random"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userDBHandler interface {
	Read(ctx context.Context, username string) (*models.User, error)
}

type sessionDBHandler interface {
	Create(ctx context.Context, session *models.Session) error
}

// HTTPHandler ...
type HTTPHandler struct {
	appCtx app.Context
	userDBHandler
	sessionDBHandler
}

// NewHTTPHandler creates a new login http handler
func NewHTTPHandler(appCtx app.Context, userDBHandler userDBHandler, sessionHandler sessionDBHandler) *HTTPHandler {
	return &HTTPHandler{
		appCtx:           appCtx,
		userDBHandler:    userDBHandler,
		sessionDBHandler: sessionHandler,
	}
}

// Post ...
func (h *HTTPHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	username, password, ok := request.BasicAuth()
	if !ok {
		log.Error().Println("failed to get username and password from authorization header")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userDBHandler.Read(ctx, username)
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	session := models.Session{
		ID:     random.LongString(),
		UserID: user.ID,
	}
	err = h.sessionDBHandler.Create(ctx, &session)
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: make secure cookie
	cookie := http.Cookie{
		Name:     h.appCtx.Cfg.SessionCookieName,
		Value:    session.ID,
		Domain:   h.appCtx.Cfg.Host,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(writer, &cookie)

	log.Info().Printf("logged in user %s with session id %s\n", user.ID, session.ID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
