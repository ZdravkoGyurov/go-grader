package account

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/app"
)

func UserLoggedIn(appContext app.Context, request *http.Request) bool {
	_, err := request.Cookie(appContext.Cfg.SessionCookieName)
	return err == nil
}
