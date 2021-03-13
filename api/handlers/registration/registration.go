package registration

import (
	"context"
	"grader/db/models"
	"grader/log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userDBHandler interface {
	Create(ctx context.Context, user *models.User) error
}

// HTTPHandler ...
type HTTPHandler struct {
	userDBHandler
}

// NewHTTPHandler creates a new registration http handler
func NewHTTPHandler(userDBHandler userDBHandler) *HTTPHandler {
	return &HTTPHandler{
		userDBHandler: userDBHandler,
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

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error().Printf("failed to hash the password: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:       uuid.NewString(),
		Username: username,
		Fullname: "fullname", // TODO: get from body
		Password: string(hash),
		Role:     "STUDENT", // TODO: add roles
		Disabled: false,
	}

	if err := h.userDBHandler.Create(ctx, &user); err != nil {
		log.Error().Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Printf("created user with id %s\n", user.ID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
