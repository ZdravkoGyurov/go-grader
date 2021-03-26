package api

import (
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
)

func StatusCode(err error) int {
	if errors.Is(err, errors.ErrInvalidInput) {
		return http.StatusBadRequest
	}
	if errors.Is(err, errors.ErrNoDocuments) {
		return http.StatusNotFound
	}
	if errors.Is(err, errors.ErrIsDuplicate) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}
