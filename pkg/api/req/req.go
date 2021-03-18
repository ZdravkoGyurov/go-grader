package req

import (
	"context"
	"net/http"
)

const (
	PermissionsKey = iota
)

func AddPermissions(r *http.Request, permissions []string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), PermissionsKey, permissions))
}

func GetPermissions(r *http.Request) ([]string, bool) {
	permissions, ok := r.Context().Value(PermissionsKey).([]string)
	return permissions, ok
}
