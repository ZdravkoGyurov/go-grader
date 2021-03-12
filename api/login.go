package api

import (
	"context"
	"grader/db/models"
	"grader/log"
	"grader/random"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userHandler interface {
	ReadUser(ctx context.Context, username string) (*models.User, error)
}

type sessionDBHandler interface {
	CreateSession(ctx context.Context, session *models.Session) error
}

// LoginHandler ...
type LoginHandler struct {
	userHandler
	sessionDBHandler
}

// NewLoginHandler creates a new login http handler
func NewLoginHandler(userHandler userHandler, sessionHandler sessionDBHandler) *LoginHandler {
	return &LoginHandler{
		userHandler:      userHandler,
		sessionDBHandler: sessionHandler,
	}
}

// Post ...
func (h *LoginHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	username, password, ok := request.BasicAuth()
	if !ok {
		log.Error().Println("failed to get username and password from authorization header")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userHandler.ReadUser(ctx, username)
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
	err = h.sessionDBHandler.CreateSession(ctx, &session)
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: make secure cookie
	cookie := http.Cookie{
		Name:     "Grader",
		Value:    session.ID,
		Domain:   "localhost",
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
