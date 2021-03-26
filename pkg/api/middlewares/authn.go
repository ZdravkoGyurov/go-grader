package middlewares

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/req"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
)

type Authentication struct {
	Controller *controller.Controller
}

func (m Authentication) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		cookie, err := request.Cookie(m.Controller.Config.SessionCookieName)
		if err != nil {
			response.SendError(writer, http.StatusUnauthorized, err)
			return
		}

		sessionID := cookie.Value
		user, err := m.Controller.GetUserBySessionID(ctx, sessionID)
		if err != nil {
			response.SendError(writer, http.StatusInternalServerError, err)
		}

		data := req.Data{
			Permissions:    user.Permissions,
			GithubUsername: user.GithubUsername,
		}
		request = req.AddRequestData(request, data)

		next.ServeHTTP(writer, request)
	})
}
