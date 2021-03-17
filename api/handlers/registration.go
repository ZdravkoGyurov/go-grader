package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

type registrationStorage interface {
	CreateUser(ctx context.Context, user *model.User) error
}

type registrationHandler struct {
	appContext app.Context
	registrationStorage
}

func (h *registrationHandler) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if _, err := request.Cookie(h.appContext.Cfg.SessionCookieName); err == nil {
		log.Error().Println(errors.New("failed to register logged in user"))
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	username, password, ok := request.BasicAuth()
	if !ok {
		log.Error().Println("failed to get username and password from authorization header")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error().Printf("failed to hash the password: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := model.User{
		ID:          uuid.NewString(),
		Username:    username,
		Fullname:    "fullname", // TODO: get from body
		Password:    string(hash),
		Permissions: []string{"STUDENT"}, // TODO: add permissions
		Disabled:    false,
	}

	if err := h.registrationStorage.CreateUser(ctx, &user); err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Printf("created user with id %s\n", user.ID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
