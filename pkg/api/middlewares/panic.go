package middlewares

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				err := errors.Newf("recovered from panic %s", r)
				response.SendError(writer, http.StatusInternalServerError, err)
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
