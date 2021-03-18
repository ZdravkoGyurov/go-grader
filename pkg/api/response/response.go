package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

func Send(writer http.ResponseWriter, status int, data interface{}, err error) {
	if err != nil {
		if status == http.StatusInternalServerError {
			respondInternalError(writer)
		} else {
			respondError(writer, status, err.Error())
		}
		log.Error().Println(err)
		return
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		respondInternalError(writer)
		log.Error().Printf("failed to marshal response %+v", data)
		return
	}

	respond(writer, status, jsonBytes)
}

func respondInternalError(writer http.ResponseWriter) {
	errorMsg := http.StatusText(http.StatusInternalServerError)
	respondError(writer, http.StatusInternalServerError, errorMsg)
}

func respondError(writer http.ResponseWriter, status int, errorMsg string) {
	respond(writer, status, []byte(fmt.Sprintf(`{"error":"%s"}`, errorMsg)))
}

func respond(writer http.ResponseWriter, status int, jsonBytes []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(jsonBytes)
	log.Info().Printf("responded with %d - %s\n", status, string(jsonBytes))
}
