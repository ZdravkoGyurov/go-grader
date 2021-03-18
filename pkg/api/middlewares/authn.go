package middlewares

import (
	"context"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/req"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/app"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

type authnStorage interface {
	ReadSession(ctx context.Context, sessionID string) (*model.Session, error)
	ReadUserByID(ctx context.Context, userID string) (*model.User, error)
}

type authnMiddleware struct {
	appContext   app.Context
	authnStorage authnStorage
}

func (m authnMiddleware) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		cookie, err := request.Cookie(m.appContext.Cfg.SessionCookieName)
		if err != nil {
			log.Error().Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value
		session, err := m.authnStorage.ReadSession(ctx, sessionID)
		if err != nil {
			response.Send(writer, http.StatusInternalServerError, nil, err)
			return
		}

		user, err := m.authnStorage.ReadUserByID(ctx, session.UserID)
		if err != nil {
			log.Error().Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		request = req.AddPermissions(request, user.Permissions)

		next.ServeHTTP(writer, request)
	})
}
