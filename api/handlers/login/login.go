package login

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db/models"
	"github.com/ZdravkoGyurov/go-grader/log"
	"github.com/ZdravkoGyurov/go-grader/random"
)

type userDBHandler interface {
	ReadByUsername(ctx context.Context, username string) (*models.User, error)
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

	_, err := request.Cookie(h.appCtx.Cfg.SessionCookieName)
	if err == nil {
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

	user, err := h.userDBHandler.ReadByUsername(ctx, username)
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
