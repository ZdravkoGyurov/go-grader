package authn

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/api/req"
	"github.com/ZdravkoGyurov/go-grader/app"
	"github.com/ZdravkoGyurov/go-grader/db/handlers/session"
	"github.com/ZdravkoGyurov/go-grader/db/handlers/user"
	"github.com/ZdravkoGyurov/go-grader/log"
)

type middleware struct {
	appCtx           app.Context
	sessionDBHandler *session.DBHandler
	userDBHandler    *user.DBHandler
}

func Middleware(appCtx app.Context, sessionDBHandler *session.DBHandler, userDBHandler *user.DBHandler) func(http.Handler) http.Handler {
	mw := &middleware{
		appCtx:           appCtx,
		sessionDBHandler: sessionDBHandler,
		userDBHandler:    userDBHandler,
	}
	return mw.authenticate
}

func (m middleware) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		cookie, err := request.Cookie(m.appCtx.Cfg.SessionCookieName)
		if err != nil {
			log.Error().Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value
		session, err := m.sessionDBHandler.Read(ctx, sessionID)
		if err != nil {
			log.Error().Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := m.userDBHandler.ReadByID(ctx, session.UserID)
		if err != nil {
			log.Error().Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		request = req.AddPermissions(request, user.Permissions)

		next.ServeHTTP(writer, request)
	})
}
