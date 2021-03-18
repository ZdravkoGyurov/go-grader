package middlewares

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				log.Error().Printf("recovered from panic: %s", err)
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
