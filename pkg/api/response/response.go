package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

func SendData(writer http.ResponseWriter, status int, data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		respondInternalError(writer)
		log.Error().Printf("failed to marshal response %+v", data)
		return
	}
	respond(writer, status, jsonBytes)
}

func SendError(writer http.ResponseWriter, status int, err error) {
	if err != nil {
		respondInternalError(writer)
		log.Error().Println("failed to return nil error")
		return
	}
	if status == http.StatusInternalServerError {
		respondInternalError(writer)
	} else {
		respondError(writer, status, err.Error())
	}
	log.Error().Println(err)
	return
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
