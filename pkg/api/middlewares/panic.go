package middlewares

import (
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				response.Send(writer, http.StatusInternalServerError, nil, fmt.Errorf("recovered from panic: %s", err))
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
