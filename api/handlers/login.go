package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/ZdravkoGyurov/go-grader/pkg/random"
)

type loginStorage interface {
	ReadUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateSession(ctx context.Context, session *model.Session) error
}

type loginHandler struct {
	appContext app.Context
	loginStorage
}

func (h *loginHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.appContext.Cfg.SessionCookieName); err == nil {
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

	user, err := h.loginStorage.ReadUserByUsername(ctx, username)
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

	session := model.Session{
		ID:     random.String(20),
		UserID: user.ID,
	}
	err = h.loginStorage.CreateSession(ctx, &session)
	if err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: make secure cookie
	cookie := http.Cookie{
		Name:     h.appContext.Cfg.SessionCookieName,
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
