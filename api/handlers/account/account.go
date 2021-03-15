package account

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/app"
)

func UserLoggedIn(appCtx app.Context, request *http.Request) bool {
	_, err := request.Cookie(appCtx.Cfg.SessionCookieName)
	return err == nil
}
